package main

import (
	"log"
	"math/rand"
	"net"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)

	listener, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Accept Error: %v\n", err)
			continue
		}

		log.Printf("Accept [%v]\n", conn.RemoteAddr())

		go doProcess(conn)
	}
}

func doProcess(conn net.Conn) {
	var rlen int
	var err error

	tcpConn := conn.(*net.TCPConn)

	defer tcpConn.Close()

	err = tcpConn.SetDeadline(time.Now().Add(5 * time.Second))

	if err != nil {
		log.Printf("[%v]: %v\n", tcpConn.RemoteAddr(), err)
		return
	}

	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	buf := make([]byte, 1024)

	rlen, err = tcpConn.Read(buf)

	if err != nil {
		log.Printf("Receive Error [%v]: %v\n", tcpConn.RemoteAddr(), err)
		return
	}

	s := string(buf[:rlen])

	log.Printf("Receive [%v]: %v\n", tcpConn.RemoteAddr(), s)

	s = "Hello! " + s

	rlen, err = tcpConn.Write([]byte(s))

	if err != nil {
		log.Printf("Send Error [%v]: %v\n", tcpConn.RemoteAddr(), err)
		return
	}

	log.Printf("Send [%v]: %v\n", tcpConn.RemoteAddr(), s)
}
