package main

import (
	"fmt"
	"net"
)

func main() {
	// TCP ports range from 1 to 65,535
	// For testing, we'll only scan ports 1 to 1024
	for i := 1; i <= 1024; i++ {
		// We can convert an int into a string in two ways in Go.
		// 1) Use the string conversion package strconv.
		// 2) Use Sprintf(format string, a ...interface{}) from
		//    the fmt package which returns a generated string.
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		// Dial takes two args: (network, address string)
		// network specifies the kind of connection: tcp, udp, unix, etc.
		// address specifies the host to which you want to connect.
		// address will be of the format "host:port"
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered (firewall).
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
