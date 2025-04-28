package main

import (
	"fmt"
)

// 4/9 Обработка ошибок
// Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления,
// но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции.
// Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error).
// Пакет main уже объявлен, а нужные пакеты уже импортированы!
func main() {
	var number1 int
	var err1 error
	_, err1 = fmt.Scan(&number1)
	var number2 int
	var err2 error
	_, err2 = fmt.Scan(&number2)

	if err1 != nil || err2 != nil || number1 == 0 || number2 == 0 {
		fmt.Println("ошибка")
	} else {
		fmt.Println(number1 / number2)
	}
}
