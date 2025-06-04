package main

import (
	"fmt"
	"net/http"
)

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello," + r.URL.Query().Get("name") + "!"))
}

// 4/9 Веб-сервера
// Напишите веб-сервер который по пути /api/user приветствует пользователя:
// Принимает и парсит параметр name и делает ответ "Hello,<name>!"
// Пример: /api/user?name=Golang
// Ответ: Hello,Golang!
// порт :9000
func main() {
	http.HandleFunc("/api/user", handler2)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
