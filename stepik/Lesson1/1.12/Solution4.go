package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 16/16 Массивы и срезы
// Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.
// Входные данные
// Сначала задано число N — количество элементов в последовательности (1≤N≤100). Далее через пробел записаны N чисел — элементы последовательности.
// Последовательность состоит из целых чисел.
// Выходные данные
// Необходимо вывести единственное число - количество положительных элементов в последовательности.
func main() {

	var result int
	var N int
	fmt.Scan(&N)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	resultArray := strings.Split(str, " ")
	for i := 0; i < N; i++ {
		resultInt, _ := strconv.Atoi(resultArray[i])
		if resultInt >= 0 {
			result++
		}
	}

	fmt.Println(result)
}
