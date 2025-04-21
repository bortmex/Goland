package main

import (
	"fmt"
)

var mapFibos = map[int]int{}

func fibonacciRecursive(n int) int {
	_, ok := mapFibos[n]
	if ok {
		return mapFibos[n-1] + mapFibos[n-2]
	} else {
		return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
	}

}

// функции ниже достаточно
func fibonacci(n int) int {

	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)

}

// 9/12 Функции
// Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
// Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
// Напишите функцию, которая по указанному натуральному n возвращает φn.
// Входные данные
// В функцию передается одно число n.
// Выходные данные
// Необходимо вывести значение φn.
func main() {

	var n int
	fmt.Scan(&n)
	if n == 1 || n == 2 {
		fmt.Println(1)
		return
	} else if n == 3 {
		fmt.Println(2)
		return
	} else if n == 4 {
		fmt.Println(3)
		return
	}
	mapFibos[1] = 1
	mapFibos[2] = 1
	mapFibos[3] = 2
	mapFibos[4] = 3
	for i := 5; ; i++ {
		recursive := fibonacciRecursive(i)
		if n > i {
			_, ok1 := mapFibos[i-4]
			if ok1 {
				delete(mapFibos, i-4)
			}
			mapFibos[i] = recursive
		} else if n == i {
			fmt.Println(recursive)
			break
		} else {
			fmt.Println(-1)
			break
		}
	}
}
