package main

import "fmt"

func main() {
	v := 5
	p := &v
	fmt.Print(*p, " ")
	changePointer(p)
	fmt.Print(*p)
}

func changePointer(p *int) {
	v := 3
	p = &v
	fmt.Println("пу пу пу")
}

//
//package main
//
//import "fmt"
//
//func main() {
//	v := 5
//	p := &v
//	fmt.Print(*p, " ")
//	changePointer(p)
//	fmt.Print(*p)
//}
//
//func changePointer(p *int) {
//	v := 333
//	*p = v
//	fmt.Println("пу пу пу")
//}

//func main() {
//	v := 5
//	p := &v
//	fmt.Print(*p, " ")
//	changePointer(&p)
//	fmt.Print(*p)
//}
//
//func changePointer(p **int) {
//	v := 3
//	*p = &v
//	fmt.Println("пу пу пу")
//}
