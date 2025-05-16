package main

import "fmt"

func convert(number int64) uint16 {
	return uint16(number)
}

// 3/14 Преобразование типов данных (map)
// Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.
func main() {
	fmt.Println(convert(123))
}
