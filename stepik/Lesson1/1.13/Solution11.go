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

// 12/14 Решение задач
// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
func main() {

	var n int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println(2)
		return
	} else if n == 2 {
		fmt.Println(3)
		return
	}
	mapFibos[0] = 1
	mapFibos[1] = 1
	mapFibos[2] = 2
	for i := 3; ; i++ {
		recursive := fibonacciRecursive(i)
		if recursive < n {
			_, ok1 := mapFibos[i-4]
			if ok1 {
				delete(mapFibos, i-4)
			}
			mapFibos[i] = recursive
		} else if recursive == n {
			fmt.Println(i + 1)
			break
		} else {
			fmt.Println(-1)
			break
		}
	}
}
