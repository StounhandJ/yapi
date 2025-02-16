package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"yapi/internal/conversation"
	"yapi/internal/env"
	"yapi/internal/glagol"
	"yapi/pkg/mdns"

	"github.com/gorilla/mux"
)

type Http struct {
	addr          string
	conversations map[string]*conversation.Conversation
}

func NewHttp(addr string) Http {
	return Http{
		addr:          addr,
		conversations: map[string]*conversation.Conversation{},
	}
}

func (h *Http) Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/", h.Write).Methods("POST")
	r.HandleFunc("/", h.Read).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Addr:         h.addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	return srv.ListenAndServe()
}

func (s *Http) Write(w http.ResponseWriter, r *http.Request) {
	var conversationStantion *conversation.Conversation
	var err error
	var msg conversation.Payload

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if conversationStantion, err = s.getConversation(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := conversationStantion.SendToDevice(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (s *Http) Read(w http.ResponseWriter, r *http.Request) {
	var conversationStantion *conversation.Conversation
	var err error

	if conversationStantion, err = s.getConversation(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(conversationStantion.ReadFromDevice())
}

func (s *Http) getConversation(r *http.Request) (*conversation.Conversation, error) {
	var err error
	var entry mdns.Entry
	var device *glagol.Device

	username, password, ok := r.BasicAuth()
	if !ok {
		return nil, errors.New("basic auth not set")
	}

	if conversationStantion, ok := s.conversations[username]; ok {
		return conversationStantion, nil
	}

	if entry, err = mdns.Discover(username, mdns.YandexServicePrefix); err != nil {
		return nil, err
	}

	client := glagol.NewClient(env.Config.GlagolUrl, username, password)
	if device, err = client.GetDevice(context.Background(), entry.IpAddr, entry.Port); err != nil {
		return nil, err
	}

	conversationStantion := conversation.NewConversation(device)

	if err = conversationStantion.Connect(context.TODO()); err != nil {
		return nil, err
	}

	s.conversations[username] = conversationStantion
	return conversationStantion, nil
}
