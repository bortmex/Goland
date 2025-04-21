package main

import "fmt"

func sumInt(intValue ...int) (count, sum int) {
	for _, elem := range intValue {
		sum += elem
		count++
	}
	return
}

// 11/12 Функции
// Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
// Пример вызова вашей функции:
// a, b := sumInt(1, 0)
// fmt.Println(a, b)
// Результат: 2, 1
func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}
