package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 15/16 Массивы и срезы
// Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
// Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
//
// Входные данные
// Сначала задано число N — количество элементов в массиве (1≤N≤100). Далее через пробел записаны N чисел — элементы массива. Массив состоит из целых чисел.
// Выходные данные
// Необходимо вывести все элементы массива с чётными индексами.
func main() {

	var N int
	fmt.Scan(&N)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	resultArray := strings.Split(str, " ")
	for i := 0; i < N; i++ {
		resultInt, _ := strconv.Atoi(resultArray[i])
		if i%2 == 0 {
			fmt.Printf("%v ", resultInt)
		}
	}
}
