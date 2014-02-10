package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8888")

	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		rlen, remote, err := conn.ReadFromUDP(buf)

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		s := string(buf[:rlen])

		log.Printf("Receive [%v]: %v\n", remote, s)

		s = "Hello! " + s

		rlen, err = conn.WriteToUDP([]byte(s), remote)

		if err != nil {
			log.Printf("Receive Error [%v]: %v\n", remote, s)
		}

		log.Printf("Send [%v]: %v\n", remote, s)
	}
}
