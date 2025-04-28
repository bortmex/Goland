package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 4/7 Решение задач
// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
// Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
func main() {
	var str string
	fmt.Scan(&str)
	arrayStr := strings.Split(str, "")

	var resultMax string

	for i := 0; i < len(arrayStr); i++ {
		number, _ := strconv.Atoi(arrayStr[i])
		pow2 := math.Pow(float64(number), 2)
		resultMax += strconv.Itoa(int(pow2))
	}

	fmt.Println(resultMax)
}
