package main

import "fmt"

// 9/14 Решение задач
// Самое большое число, кратное 7
// Найдите самое большее число на отрезке от a до b, кратное 7 .
func main() {

	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	minus := m % 7
	if minus == 0 {
		fmt.Println(m)
	} else {
		if minus < 0 {
			minus = 7 + minus
		}
		if m-minus >= n {
			fmt.Println(m - minus)
		} else {
			fmt.Println("NO")
		}
	}
}
