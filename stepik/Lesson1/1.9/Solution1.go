package main

import (
	"fmt"
	"strconv"
)

// 5/9 По данному трехзначному числу определите, все ли его цифры различны.
// Формат входных данных
// На вход подается одно натуральное трехзначное число.
// Формат выходных данных
// Выведите "YES", если все цифры числа различны, в противном случае - "NO".
func main() {
	var number int
	fmt.Scan(&number)
	m := make(map[int]int)
	str := strconv.Itoa(number)

	for i := 0; i < len(str); i++ {
		local1, _ := strconv.Atoi(string(str[i]))
		if m[local1] == 0 {
			m[local1] = 1
		} else {
			m[local1] = m[local1] + 1
		}
	}

	for _, value := range m {
		if value > 1 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
