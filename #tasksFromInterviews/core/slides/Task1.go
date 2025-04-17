package main

import "fmt"

func foo(src []int) {
	src = append(src, 15)
}

func main() {
	arr := []int{1, 2, 3}
	src := arr[:1]

	foo(src)

	fmt.Println("arr:", arr)
	fmt.Println("src:", src)
	fmt.Println("len(src):", len(src))
	fmt.Println("cap(src):", cap(src))
	fmt.Println("srcNEW len 3:", src[:3])
}
