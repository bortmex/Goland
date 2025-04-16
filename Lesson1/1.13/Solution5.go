package main

import (
	"fmt"
)

// 6/13 Решение задач
// По данным числам, определите количество чисел, которые равны нулю.
func main() {

	var N int
	var a int
	fmt.Scan(&N)
	var result int
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		if a == 0 {
			result++
		}
	}

	fmt.Println(result)
}
