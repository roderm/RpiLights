package server

import (
	"fmt"
	"log"
	"net"
	"playground/light"
	"strconv"
	"strings"
)

type TelnetSever struct {
	Light *light.RpiLight
}

func (t *TelnetSever) Serve(port int) error {
	host := fmt.Sprintf("%s:%d", getOutboundIP().String(), port)
	fmt.Println("Start on " + host)

	l, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	defer l.Close()
	fmt.Println("Listening on " + host)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			break
		}
		// Handle connections in a new goroutine.
		go t.handleRequest(conn)
	}
	return nil
}

func (t *TelnetSever) handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	cStrs := string(buf[:])
	command := strings.Split(cStrs, " ")
	switch command[0] {
	case "on":
		go t.Light.On()
	case "off":
		go t.Light.Off()
	case "color":
		if len(command) >= 4 {
			r, err := strconv.Atoi(command[1])
			g, err := strconv.Atoi(command[2])
			b, err := strconv.Atoi(command[3])
			if err != nil {
				conn.Write([]byte("Invalid colors" + err.Error()))
				return
			}
			go t.Light.SetColors(light.ColorScheme{
				Red:   uint8(r),
				Green: uint8(g),
				Blue:  uint8(b)})
		} else {
			conn.Write([]byte("3 coulours are used"))
		}
	case "bright":
		if len(command) >= 2 {
			b, err := strconv.Atoi(command[1])
			if err != nil {
				conn.Write([]byte("Invalid bright" + err.Error()))
				return
			}
			if b > 100 {
				b = 100
			}
			if b < 100 {
				b = 1
			}
			go t.Light.DimTo(b)
		} else {
			conn.Write([]byte("Brightness between 1 and 100 used"))
		}
	default:
		conn.Write([]byte("No command found for: " + command[0]))
	}

	// Close the connection when you're done with it.
	conn.Close()
}
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
