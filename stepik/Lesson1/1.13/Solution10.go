package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 11/14 Решение задач
// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
func main() {

	var N int
	fmt.Scan(&N)
	var array []string
	for i := 0; ; i++ {
		result := math.Pow(2, float64(i))
		resultInt := int(result)
		if resultInt <= N {
			resultStr := strconv.Itoa(resultInt)
			array = append(array, resultStr)
		} else {
			break
		}
	}

	justString := strings.Join(array, " ")
	fmt.Println(justString)
}
