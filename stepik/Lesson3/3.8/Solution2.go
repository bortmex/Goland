package main

import "fmt"

// 10/10 Параллелизм ч.1
// Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет
// значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.
// Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки,
// во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения,
// которые не повторяются подряд. Не забудьте закрыть канал ;)
// Функция должна называться removeDuplicates()
func removeDuplicates(inputStream chan string, outputStream chan string) {
	var check string
	for v := range inputStream {
		if check != v {
			outputStream <- v
			check = v
		}
	}

	close(outputStream)
}

func main() {
	inputStream := make(chan string, 5)
	outputStream := make(chan string, 5)
	go removeDuplicates(inputStream, outputStream)
	for _, v := range []string{"a", "a", "b", "b", "c", "c", "d", "d", "e"} {
		inputStream <- v
	}
	close(inputStream)
	for result := range outputStream {
		fmt.Println(result)
	}
}
