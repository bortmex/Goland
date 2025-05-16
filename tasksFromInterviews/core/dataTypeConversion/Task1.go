package main

import (
	"fmt"
	"math"
)

func main() {

	var index int8 = 15
	var bigIndex int32
	bigIndex = int32((index)) + 1

	fmt.Println(index)
	fmt.Println(bigIndex)
	fmt.Printf("%T \n", index)
	fmt.Printf("%T \n", bigIndex)

	fmt.Println(math.MaxInt32)
}
