package main

import (
	"fmt"
	"math"
)

// 2/10 Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10.
// Квадрат каждого числа должен выводится в новой строке.
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(math.Pow(float64(i), 2))
	}
}