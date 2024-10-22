package main

import (
	"fmt"
	"math"
)

// 3/11 Форматированный вывод
func main() {
	var n float64
	fmt.Scan(&n)
	result := math.Pow(n, 2)
	if n <= 0.0 {
		fmt.Printf("число %2.2f не подходит", n)
	} else if n > 10000 {
		fmt.Printf("%e", n)
	} else {
		fmt.Printf("%.4f", result)
	}
}
