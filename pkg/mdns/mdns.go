package mdns

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/mdns"
)

type Entry struct {
	Discovered bool
	IpAddr     string
	Port       string
}

func Discover(deviceId, service string) (entry Entry, err error) {
	entriesCh := make(chan *mdns.ServiceEntry)
	defer close(entriesCh)

	go func(entry *Entry) {
		for entryCh := range entriesCh {
			if deviceId == getDeviceId(entryCh) {
				entry.IpAddr = entryCh.AddrV4.String()
				entry.Port = strconv.Itoa(entryCh.Port)
				entry.Discovered = true
				log.Println("Found device on: " + entry.IpAddr)
				return
			}
		}
	}(&entry)
	err = mdns.Lookup(service, entriesCh)
	if err != nil {
		return
	}
	if !entry.Discovered {
		err = errors.New("mdns: No device found")
	}

	return
}

func getDeviceId(entry *mdns.ServiceEntry) (deviceId string) {
	for _, field := range entry.InfoFields {
		entryData := strings.Split(field, "=")
		if len(entryData) == 2 && entryData[0] == "deviceId" {
			deviceId = entryData[1]
			return
		}
	}
	return
}
