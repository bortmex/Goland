package main

import (
	"fmt"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

// 2/9 Веб-сервера
// Напишите веб сервер, который по пути /get отдает текст "Hello, web!".
// Порт должен быть :8080.
func main() {
	http.HandleFunc("/", handler1)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
