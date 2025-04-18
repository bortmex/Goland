package main

import (
	"fmt"
)

// 5/14 Решение задач
// Даны два числа. Найти их среднее арифметическое.
func main() {

	var n1 int
	var n2 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Println((float64(n1) + float64(n2)) / 2)
}
