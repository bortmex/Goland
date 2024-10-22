package main

import (
	"fmt"
)

// 3/10 Требуется написать программу, при выполнении которой с клавиатуры считываются
// два натуральных числа A и B (каждое не более 100, A < B).
// Вывести сумму всех чисел от A до B  включительно.
func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	var sum int

	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(sum)
}
