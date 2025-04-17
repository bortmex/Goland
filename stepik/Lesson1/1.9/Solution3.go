package main

import (
	"fmt"
	"strconv"
)

// 7/9 Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
// Формат входных данных
// На вход подается номер билета - одно шестизначное  число.
// Формат выходных данных
// Выведите "YES", если билет счастливый, в противном случае - "NO".
func main() {
	var number int
	fmt.Scan(&number)
	var strs1 = [3]int{}
	var strs2 = [3]int{}
	str := strconv.Itoa(number)

	for i := 0; i < 3; i++ {
		strs1[i], _ = strconv.Atoi(string(str[i]))
	}
	for i := 3; i < 6; i++ {
		strs2[i-3], _ = strconv.Atoi(string(str[i]))
	}

	var result1 int
	var result2 int

	for i := 0; i < 3; i++ {
		result1 += strs1[i]
	}
	for i := 0; i < 3; i++ {
		result2 += strs2[i]
	}

	if result1 == result2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
