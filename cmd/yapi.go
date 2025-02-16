package main

import (
	"log"
	"yapi/internal/env"
	"yapi/internal/server"
)

func main() {

	srv := server.NewHttp(env.Config.HttpHost)
	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}
}
