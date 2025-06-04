package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int

func handler3(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(counter)))
	case "POST":
		count := r.FormValue("count")
		countNumber, err := strconv.Atoi(count)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
		} else {
			counter += countNumber
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(strconv.Itoa(counter)))
		}
	}
}

// 7/9 Веб-сервера
// GET: возвращает счетчик
// POST: увеличивает ваш счетчик на значение (с ключом "count") которое вы получаете из формы,
// но если пришло НЕ число то нужно ответить клиенту: "это не число" со статусом http.StatusBadRequest (400).
func main() {
	http.HandleFunc("/", handler3)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
