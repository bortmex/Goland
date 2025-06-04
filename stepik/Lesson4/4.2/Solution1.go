package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// 9/10 Работа с сетью
// Сделайте HTTP запрос на сервер по пути http://127.0.0.1:5555/get и напечатайте ответ сервера (только тело).
func main() {
	resp, err := http.Get("http://127.0.0.1:5555/get")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%s", data)
}
