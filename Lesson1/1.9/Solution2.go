package main

import (
	"fmt"
	"strconv"
)

// 6/9 Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.
// Формат выходных данных
// Выведите одно целое число - первую цифру заданного числа.
func main() {
	var number int
	fmt.Scan(&number)
	str := strconv.Itoa(number)
	fmt.Println(string(str[0]))
}
