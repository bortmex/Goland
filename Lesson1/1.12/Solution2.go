package main

import "fmt"

// 14/16 Массивы и срезы
// На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.
// Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
func main() {

	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	result := array[0]

	for i := 1; i < 5; i++ {
		if result < array[i] {
			result = array[i]
		}
	}

	fmt.Println(result)
}
