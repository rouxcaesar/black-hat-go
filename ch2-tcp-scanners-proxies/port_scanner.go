package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "canme.nmap.org:80")
	if err != nil {
		fmt.Println("Connection successful")
	}
}
