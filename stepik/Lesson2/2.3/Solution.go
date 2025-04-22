package main

import "fmt"

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

// 5/8 Указатели
// Напишите функцию, которая умножает значения на которые
// ссылаются два указателя и выводит получившееся произведение в консоль.
// Ввод данных уже написан за вас.
func main() {
	number1 := 2
	number2 := 2
	test(&number1, &number2)
}
