package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

func Work(wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()
	defer conn.Close()

	buf := make([]byte, 1024)
	for i := 0; i < 3; i++ {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		str := strings.ToUpper(string(buf[:n]))
		fmt.Println(str)
	}
}

// 4/10 Работа с сетью
// Подключитесь к адресу 127.0.0.1:8081 по протоколу TCP, считайте от сервера 3 сообщения,
// и выведите их в верхнем регистре. Рекомендуем использовать буфер в 1024 байта.
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Work(wg, conn)
	wg.Wait()
}
