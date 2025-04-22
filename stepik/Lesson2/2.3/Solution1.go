package main

import "fmt"

func test1(x1 *int, x2 *int) {
	x3 := *x1
	*x1 = *x2
	*x2 = x3
	fmt.Println(*x1, *x2)
}

// 6/8 Указатели
// Поменяйте местами значения переменных на которые ссылаются указатели.
// После этого переменные нужно вывести.
func main() {
	number1 := 2
	number2 := 4
	test1(&number1, &number2)
	fmt.Println(number1, number2)
}
