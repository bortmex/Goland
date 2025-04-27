package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 9/12 Строки
// Даются две строки X и S. Нужно найти и вывести индекс первого вхождения подстроки S в строке X.
// Если подстроки S нет в строке X - вывести -1
func main() {
	reader := bufio.NewReader(os.Stdin)
	X, _ := reader.ReadString('\n')
	X = strings.ReplaceAll(X, "\n", "")
	S, _ := reader.ReadString('\n')
	S = strings.ReplaceAll(S, "\n", "")

	index := strings.Index(X, S)
	if index >= 0 {
		fmt.Println(index)
	} else {
		fmt.Println(-1)
	}
}
