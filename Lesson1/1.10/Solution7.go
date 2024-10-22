package main

import "fmt"

// 8/10 Вклад в банке составляет x рублей.
// Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
// Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
// Входные данные
// Программа получает на вход три натуральных числа: x, p, y.
// Выходные данные
// Программа должна вывести одно целое число.
func main() {
	var x int
	var p int
	var y int

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	j := 0
	for i := x; i <= y; {
		j++
		i = i + i*p/100
		x = i
	}

	fmt.Println(j)
}