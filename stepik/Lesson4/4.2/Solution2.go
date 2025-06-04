package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// 10/10 Работа с сетью
// Считайте с консоли две переменные, сначала name, затем age.
// Сделайте HTTP запрос на http://127.0.0.1:8080/hello передав name и age в query параметрах, затем напечатайте ответ сервера (только тело).
func main() {

	var N1 string
	fmt.Scan(&N1)
	var N2 string
	fmt.Scan(&N2)
	baseURL := "http://127.0.0.1:8080/hello"
	params := url.Values{}
	params.Add("name", N1)
	params.Add("age", N2)

	fullURL := baseURL + "?" + params.Encode()

	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Ошибка при отправке GET-запроса:", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%s", data)
}
