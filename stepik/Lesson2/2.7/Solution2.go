package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 4/7 Решение задач
// Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
func main() {
	var str string
	fmt.Scan(&str)
	length := utf8.RuneCountInString(str)
	if length < 1 || length > 1000 {
		return
	}
	arrayStr := strings.Split(str, "")

	var resultMax int

	for i := 0; i < len(arrayStr); i++ {
		number, _ := strconv.Atoi(arrayStr[i])
		if resultMax < number {
			resultMax = number
		}
	}

	fmt.Println(resultMax)
}
