package main

import (
	"fmt"
)

// 7/14 Решение задач
// Найдите количество минимальных элементов в последовательности.
func main() {

	var N int
	var a int
	fmt.Scan(&N)
	fmt.Scan(&a)
	result1 := a
	var array []int
	array = append(array, a)
	for i := 1; i < N; i++ {
		fmt.Scan(&a)
		if a < result1 {
			result1 = a
		}
		array = append(array, a)
	}

	var result int
	for i := 0; i < len(array); i++ {
		if array[i] == result1 {
			result++
		}
	}

	fmt.Println(result)
}
