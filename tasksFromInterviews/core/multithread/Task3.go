package main

import "fmt"

func main() {

	c := make(chan int, 5)
	for i := 0; i < 20; i++ {
		go func(i int) {
			c <- i
			fmt.Println(i)
		}(i)
	}
	fmt.Print(<-c + 1000)
}
