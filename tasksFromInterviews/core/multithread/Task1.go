package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	d := 42

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("func():", d*2)
	}()

	go func(d int) {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("func(d int):", d*2)
	}(d)

	d = 1

	wg.Wait()
}
