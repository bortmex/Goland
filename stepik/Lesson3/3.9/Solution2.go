package main

import (
	"fmt"
)

func calculator1(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	var val int
	go func() {
		defer close(ch)
		select {
		case val = <-firstChan:
			ch <- val * val
		case val = <-secondChan:
			ch <- val * 3
		case <-stopChan:
			return
		}
	}()
	return ch
}

// 13/16 Параллелизм ч.2
// Вам необходимо написать функцию calculator следующего вида:
// func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
// Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
// в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный)
// канал вы должны отправить квадрат аргумента.
// в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный)
// канал вы должны отправить результат умножения аргумента на 3.
// в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
// Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит
// всего одно значение в один из каналов - получили значение, обработали его, завершили работу.
// После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете,
// то превысите предельное время работы.
func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	resultChan := calculator1(firstChan, secondChan, stopChan)

	// Раскомментируйте нужный канал для проверки поведения
	firstChan <- 2
	//secondChan <- 3
	//close(stopChan)

	for v := range resultChan {
		fmt.Println(v)
	}

	close(firstChan)
	close(secondChan)
}
