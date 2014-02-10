package main

import (
	"log"
	"net"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)

	const max = 10

	ch := make(chan int, max)

	for i := 0; i < max; i++ {
		go exec(ch, i)
	}

	for i := 0; i < max; i++ {
		<-ch
	}
}

func exec(ch chan<- int, pos int) {
	var rlen int
	var err error

	defer func() {
		ch <- pos
	}()

	remote, err := net.ResolveUDPAddr("udp", "localhost:8888")

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	conn, err := net.DialUDP("udp", nil, remote)

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	log.Printf("Connect[%d]: %v\n", pos, remote)

	conn.SetDeadline(time.Now().Add(5 * time.Second))

	defer conn.Close()

	s := "user" + strconv.Itoa(pos)

	rlen, err = conn.Write([]byte(s))

	if err != nil {
		log.Printf("Send Error: %v\n", err)
		return
	}

	log.Printf("Send[%d]: %v\n", pos, s)

	buf := make([]byte, 1024)

	rlen, err = conn.Read(buf)

	if err != nil {
		log.Printf("Receive Error: %v\n", err)
		return
	}

	log.Printf("Receive[%d]: %v\n", pos, string(buf[:rlen]))
}
