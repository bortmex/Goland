package main

import "fmt"

// 13/16 Массивы и срезы
// Напишите программу, принимающая на вход число N(N≥4), а затем N целых чисел-элементов среза.
// На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
func main() {

	var N int
	var a int
	fmt.Scan(&N)
	var result []int
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		result = append(result, a)
	}

	fmt.Println(result[4])
}
