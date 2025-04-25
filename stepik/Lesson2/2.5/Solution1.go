package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// 8/12 Строки
// На вход подается строка. Нужно определить, является ли она палиндромом.
// Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
// Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")

	len := utf8.RuneCountInString(text)
	if len < 2 {
		fmt.Println("Палиндром")
		return
	}
	splitText := strings.Split(text, "")

	for i := 0; i < len/2; i++ {
		if splitText[i] != splitText[len-1-i] {
			fmt.Println("Нет")
			return
		}
	}

	fmt.Println("Палиндром")
}
