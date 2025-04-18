package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 14/14 Решение задач
// Из натурального числа удалить заданную цифру.
func main() {

	var number int
	var N int
	fmt.Scan(&number)
	fmt.Scan(&N)
	strNumber := strconv.Itoa(number)
	strN := strconv.Itoa(N)
	strResultNumber := strings.Replace(strNumber, strN, "", -1)
	fmt.Println(strResultNumber)
}
