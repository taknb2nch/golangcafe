package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)

	remote, err := net.ResolveUDPAddr("udp", "localhost:8888")

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	conn, err := net.DialUDP("udp", nil, remote)

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	log.Printf("Connect: %v\n", remote)

	defer conn.Close()

	ch := make(chan int, 2)

	go read(conn, ch)
	go send(conn, ch)

	for i := 0; i < 2; i++ {
		<-ch
	}
}

func read(conn *net.UDPConn, ch chan<- int) {
	var rlen int
	var err error

	defer func() { ch <- 1 }()

	for {
		buf := make([]byte, 1024)

		if rlen, err = conn.Read(buf); err != nil {
			log.Printf("Receive Error: %v\n", err)
			return
		}

		log.Printf("Receive: %v\n", string(buf[:rlen]))
	}
}

func send(conn *net.UDPConn, ch chan<- int) {
	var err error

	buf := make([]byte, 1024)

	defer func() { ch <- 1 }()

	r := bufio.NewReader(os.Stdin)

	for {
		if buf, _, err = r.ReadLine(); err != nil {
			return
		}

		_, err = conn.Write(buf)

		if err != nil {
			log.Printf("Send Error: %v\n", err)
			return
		}
	}
}
