package main

import (
	"fmt"
	"time"
)

func myFunc(i int) {
	fmt.Println("Я родился ", i)
}

// 8/10 Параллелизм ч.1
func main() {

	for i := 1; i < 1000000; i++ {
		go myFunc(i)
	}
	time.Sleep(1 * time.Second)
}
