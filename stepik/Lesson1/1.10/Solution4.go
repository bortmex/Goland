package main

import (
	"fmt"
)

// 6/10 Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
// Входные данные
// Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.
// Выходные данные
// Вывести первое число от 1 до n включительно, кратное c, но НЕ кратное d.
// Если такого числа нет - выводить ничего не нужно.
func main() {
	var a int
	var max int
	var countMax int

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a == max {
			countMax++
		} else if a > max {
			max = a
			countMax = 1
		}
	}
	fmt.Println(countMax)
}
