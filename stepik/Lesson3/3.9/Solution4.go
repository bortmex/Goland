package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		array := make([]int, n)
		wg := new(sync.WaitGroup)

		for i := 0; i < n; i++ {
			wg.Add(1)
			x1 := <-in1
			x2 := <-in2

			go func(in1 <-chan int, in2 <-chan int, i int, wg *sync.WaitGroup) {
				defer wg.Done()

				r1 := make(chan int)
				r2 := make(chan int)

				defer close(r1)
				defer close(r2)

				go calc(fn, x1, r1)
				go calc(fn, x2, r2)

				array[i] = <-r1 + <-r2
			}(in1, in2, i, wg)
		}

		wg.Wait()

		for i := 0; i < n; i++ {
			out <- array[i]
		}
	}()
}

func calc(fn func(int) int, x int, res chan<- int) {
	res <- fn(x)
}

// 15/16 Параллелизм ч.2
// Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int,
// in2 <- chan int, out chan<- int, n int).
// Описание ее работы:
// n раз сделать следующее
// прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
// вычислить f(x1) + f(x2)
// записать полученное значение в out
// Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.
// Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.
// Формат ввода:
// количество итераций передается через аргумент n.
// целые числа подаются через аргументы-каналы in1 и in2.
// функция для обработки чисел перед сложением передается через аргумент fn.
// Формат вывода:
// канал для вывода результатов передается через аргумент out.

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * x
	}
	fmt.Println("\nСоздаю каналы, начитаю работу.")
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N)
	for i := 0; i < N; i++ {
		in1 <- i
		in2 <- (i + N)
		fmt.Printf("%v) В канал 1 послано: %v, В канал 2 послано: %v.\n", i+1, i, (i + N))
	}
	fmt.Print("Функция fn с разной паузой возводит в квадрат значения из каналов.\n",
		"Итоговый результат fn(x1) + fn(x2)\n",
		"Барабанная дробь...\n")

	calcFail := false
	timeFail := false
	pass := "OK"
	for i := 0; i < N; i++ {
		c := <-out
		result := i*i + (i+N)*(i+N)
		if c != result {
			calcFail = true
			pass = "Failed."
		}
		fmt.Printf("%v) Результат: %v. Правильный %v. -> %v\n", i+1, c, result, pass)
	}
	if calcFail {
		fmt.Println("Failed. Результат выполнения не соответствует ожидаемому.")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		timeFail = true
		fmt.Println("Failed. Время превышено.")
	}
	if !calcFail && !timeFail {
		fmt.Println("Passed. OK. Ты сделал это!\nВремя работы: ", duration)
	} else {
		fmt.Println("Пока неверно. Упорства вам не занимать, попробуете еще раз?\nВремя выполнения: ", duration)
	}
}
