package main

import "fmt"

type MyFirstStruct struct {
	x, y float64
	z    string
}

func f1(c *MyFirstStruct) {
	c.x = 100
	c.y = 100
	c.z = "Карамба"
	fmt.Println("f1 ", c) // f1  &{100 100 Карамба}
}

func main() {
	c1 := MyFirstStruct{}
	c2 := new(MyFirstStruct)

	f1(&c1)
	f1(c2)

	fmt.Println("main c1 ", c1)  // main c1  {100 100 Карамба}
	fmt.Println("main c2 ", *c2) // main c2  {100 100 Карамба}
}
