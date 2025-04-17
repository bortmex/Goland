package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 1/13 Решение задач
// Дано трехзначное число. Найдите сумму его цифр.
func main() {

	var n int
	var result int
	fmt.Scan(&n)
	str := strconv.Itoa(n)
	resultArray := strings.Split(str, "")
	for i := 0; i < len(resultArray); i++ {
		number, _ := strconv.Atoi(resultArray[i])
		result += number
	}
	fmt.Println(result)
}
