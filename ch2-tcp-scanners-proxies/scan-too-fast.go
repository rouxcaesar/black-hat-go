package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// TCP ports range from 1 to 65,535
	// For testing, we'll only scan ports 1 to 1024
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			// We can convert an int into a string in two ways in Go.
			// 1) Use the string conversion package strconv.
			// 2) Use Sprintf(format string, a ...interface{}) from
			//    the fmt package which returns a generated string.
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			//address := fmt.Sprintf("127.0.0.1:%d", j)
			// Dial takes two args: (network, address string)
			// network specifies the kind of connection: tcp, udp, unix, etc.
			// address specifies the host to which you want to connect.
			// address will be of the format "host:port"
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed or filtered (firewall).
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}
