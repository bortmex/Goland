package main

import (
	"fmt"
	"strconv"
	"strings"
)

func adding(str1 string, str2 string) int64 {
	str1 = filterStr1(str1)
	str2 = filterStr1(str2)
	number1, _ := strconv.ParseInt(str1, 10, 64)
	number2, _ := strconv.ParseInt(str2, 10, 64)
	return number1 + number2
}

func filterStr1(str string) string {
	strArray := strings.Split(str, "")
	var resultStr string
	for _, symbol := range strArray {
		if strings.Contains("0123456789", symbol) {
			resultStr += symbol
		}
	}
	return resultStr
}

// 13/14 Преобразование типов данных (map)
// Представьте что вы работаете в большой компании где используется модульная архитектура.
// Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
// Вы же пишете функцию которая принимает в качестве параметров две переменных типа string ,
// а возвращает число типа int64 которое нужно получить сложением этих строк.
// Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше.
// Поэтому он решил подшутить над вами и подсунул вам свинью.
// Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.
// Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа.
// Под мусором имеются ввиду лишние символы и спец знаки.
// Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes.
func main() {
	str1 := "%^80"
	str2 := "hhhhh20&&&&nd"
	fmt.Println(adding(str1, str2))
}
