package main

import "fmt"

// 5/16 Массивы и срезы
func main() {

	workArray := [10]uint8{}
	//var workArray = [16]uint8{99, 151, 137, 71, 117, 187, 20, 93, 187, 67, 1, 2, 3, 5, 7, 8}

	//workArray[workArray[10]], workArray[workArray[11]] = workArray[workArray[11]], workArray[workArray[10]]
	//workArray[workArray[12]], workArray[workArray[13]] = workArray[workArray[13]], workArray[workArray[12]]
	//workArray[workArray[14]], workArray[workArray[15]] = workArray[workArray[15]], workArray[workArray[14]]

	var n uint8
	var y uint8

	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		workArray[i] = n
	}

	for i := 0; i < 3; i++ {
		fmt.Scan(&n)
		fmt.Scan(&y)
		workArray[n], workArray[y] = workArray[y], workArray[n]
	}

	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}
