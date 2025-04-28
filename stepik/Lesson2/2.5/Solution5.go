package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 12/12 Строки
// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
// Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль.
// Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.ReplaceAll(str, "\n", "")

	match, _ := regexp.MatchString("^[a-zA-Z0-9]{5,}$", str)
	if match {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
