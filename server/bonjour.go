package server

import (
	"github.com/hashicorp/mdns"
	"os"
)

func AdvertiseBonjour(info []string, port int) {
	// Setup our service export
	host, _ := os.Hostname()
	go func() {
		service, err := mdns.NewMDNSService(host, "_rpilight._tcp", "local.", "raspberrypi.", port, nil, info)
		if err != nil {
			panic(err)
		}
		// Create the mDNS server, defer shutdown
		server, err := mdns.NewServer(&mdns.Config{Zone: service})
		if err != nil {
			panic(err)
		}
		defer server.Shutdown()
	}()
}
