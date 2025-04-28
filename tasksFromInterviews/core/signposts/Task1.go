package main

import "fmt"

func main() {
	a := 200
	b := &a
	*b++
	c := &b
	**c++
	d := &c
	***d++
	e := &d
	****e++
	f := &e
	*****f++
	g := &f
	******g++
	fmt.Println(a) //206
}
