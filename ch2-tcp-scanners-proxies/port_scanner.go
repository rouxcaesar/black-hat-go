package main

import (
	"fmt"
	"net"
)

func main() {
	// Dial takes two args: (network, address string)
	// network specifies the kind of connection: tcp, udp, unix, etc.
	// address specifies the host to which you want to connect.
	// address will be of the format "host:port"
	_, err := net.Dial("tcp", "canme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}
