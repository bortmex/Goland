package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// 7/12 Строки
// На вход подается строка. Нужно определить, является ли она правильной или нет.
// Правильная строка начинается с заглавной буквы и заканчивается точкой.
// Если строка правильная - вывести Right иначе - вывести Wrong
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")

	if utf8.RuneCountInString(text) < 1 {
		fmt.Println("Wrong")
		return
	}
	splitText := strings.Split(text, "")
	firstLetterText := splitText[0]
	upperFirstLetterText := strings.ToUpper(firstLetterText)
	if firstLetterText == upperFirstLetterText && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
