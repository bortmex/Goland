package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 2/13 Решение задач
// Дано трехзначное число. Переверните его, а затем выведите.
func main() {

	var n int
	var result string
	fmt.Scan(&n)
	str := strconv.Itoa(n)
	resultArray := strings.Split(str, "")
	for i := len(resultArray) - 1; i >= 0; i-- {
		result += resultArray[i]
	}
	fmt.Println(result)
}
