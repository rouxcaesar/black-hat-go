package main

import "net"

func main() {
	_, err := net.Dial("tcp", "canme.nmap.org:80")
}
