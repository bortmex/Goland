package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {

	c := make(map[int]int)

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			c[i] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(len(c))
}
