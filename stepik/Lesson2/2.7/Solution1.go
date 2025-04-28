package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 3/7 Решение задач
// Дана строка, содержащая только английские буквы (большие и маленькие).
// Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
// Входные данные
// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.
// Выходные данные
// Вывести строку, которая получится после добавления символов '*'.
func main() {
	var str string
	fmt.Scan(&str)
	length := utf8.RuneCountInString(str)
	if length < 1 || length > 1000 {
		return
	}
	arrayStr := strings.Split(str, "")

	var result string

	for i := 0; i < len(arrayStr)-1; i++ {
		result += arrayStr[i] + "*"
	}
	result += arrayStr[len(arrayStr)-1]

	fmt.Println(result)
}
