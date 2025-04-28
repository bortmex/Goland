package main

import "fmt"

func foo(src *[]int) {
	*src = append(*src, 15)
}

func main() {
	arr := []int{1, 2, 3}
	src := arr[:1]

	fmt.Println("arr:", arr)
	fmt.Println("src:", src)

	foo(&src)

	fmt.Println("arr:", arr)
	fmt.Println("src:", src)
}
