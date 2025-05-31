package main

import "sync"

func work() {
}

// 5/16 Параллелизм ч.2
// Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать
// функцию work() 10 раз и дождаться результатов выполнения вызванных функций.
func main() {

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			work()
		}(wg)
	}

	wg.Wait()
}
