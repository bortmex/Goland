package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 11/12 Строки
// Дается строка. Нужно удалить все символы,
// которые встречаются более одного раза и вывести получившуюся строку
// Sample Input:
//
//	zaabcbd
//
// Sample Output:
//
//	zcd
func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	split := strings.Split(str, "")

	mapLetters := make(map[string]int)

	for i := 0; i < len(split); i++ {
		str := split[i]
		_, ok := mapLetters[str]
		if ok {
			mapLetters[str] = mapLetters[str] + 1
		} else {
			mapLetters[str] = 1
		}
	}

	for i := range mapLetters {
		if mapLetters[i] == 1 {
			delete(mapLetters, i)
		}
	}

	for i := range mapLetters {
		str = strings.ReplaceAll(str, i, "")
	}

	fmt.Println(str)
}
