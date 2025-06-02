package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8081")

	if err != nil {

		log.Println(err)

	}

	conn, err := listener.Accept()

	if err != nil {

		log.Println(err)

	}

	defer conn.Close()

	conn.Write([]byte("message1"))

	time.Sleep(10)

	conn.Write([]byte("MesSaGe2"))

	time.Sleep(10)

	conn.Write([]byte("MESSAGE3"))
}
