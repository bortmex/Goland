package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// 10/12 Строки
//
//	На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
//
// Sample Input:
//
//	ihgewlqlkot
//
// Sample Output:
//
//	hello
func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	split := strings.Split(str, "")

	var buffer bytes.Buffer

	for i := 1; i < len(split); i += 2 {
		buffer.WriteString(split[i])
	}

	fmt.Println(buffer.String())
}
