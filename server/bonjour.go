package server

import (
	"github.com/hashicorp/mdns"
	"os"
)

func AdvertiseBonjour(info []string, port int) {
	// Setup our service export
	host, _ := os.Hostname()
	service, _ := mdns.NewMDNSService(host, "_rpilight._tcp", "", "", port, nil, info)

	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()
}
