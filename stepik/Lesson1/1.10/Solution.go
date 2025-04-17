package main

import "fmt"

// 1/10 считываем числа пока не будет введен 0
func main() {
	var n int
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)
	}
}
