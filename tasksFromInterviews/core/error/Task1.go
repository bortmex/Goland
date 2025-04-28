package main

import "fmt"

func main() {
	defer finish()                                // Этот вызов отложен
	defer fmt.Println("Program has been started") // Этот вызов отложен
	panic("Ошибка")
	fmt.Println("Program is working") // Не выведется
}

func finish() {
	fmt.Println("Program has been finished")
}
