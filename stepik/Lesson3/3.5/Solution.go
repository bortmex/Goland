package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// 9/14 Работа с файлами
// Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100,
// каждое число подается на стандартный ввод с новой строки (количество чисел не известно).
// Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	result := 0
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err != nil {
			break
		}
		result += number
	}

	io.WriteString(os.Stdout, strconv.Itoa(result))
}
