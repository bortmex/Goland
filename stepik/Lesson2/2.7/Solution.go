package main

import (
	"fmt"
	"math"
)

// 2/7 Решение задач
// На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
func main() {
	var number1 int
	fmt.Scan(&number1)
	var number2 int
	fmt.Scan(&number2)

	fmt.Println(math.Sqrt(math.Pow(float64(number1), 2) + math.Pow(float64(number2), 2)))
}
