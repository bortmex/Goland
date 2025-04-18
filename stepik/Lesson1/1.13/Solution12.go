package main

import (
	"fmt"
	"strconv"
)

// 13/14 Решение задач
// Дано натуральное число N. Выведите его представление в двоичном виде.
func main() {

	var N int64
	fmt.Scan(&N)
	resultInt2 := strconv.FormatInt(N, 2)
	fmt.Println(resultInt2)
}
