package main

import "fmt"

// 12/16 Первая программа
// По данному целому числу, найдите его квадрат.
// Формат входных данных
// На вход дается одно целое число.
// Формат выходных данных
// Программа должна вывести квадрат данного числа.
func main() {
	var a int
	fmt.Scan(&a) // считаем переменную 'a' с консоли

	fmt.Println(a * a)
}
