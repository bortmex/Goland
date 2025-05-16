package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func filterStr2(str string) string {
	strArray := strings.Split(str, "")
	var resultStr string
	for _, symbol := range strArray {
		if strings.Contains(",0123456789", symbol) {
			resultStr += symbol
		}
	}
	return resultStr
}

// 14/14 Преобразование типов данных (map)
// В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела,
// кроме того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV,
// где в качестве разделителя используется символ ";".
// На стандартный ввод вы получаете 2 таких вещественных числа,
// в качестве результата требуется вывести частное от деления первого числа на второе с точностью до четырех знаков
// после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).
// Sample Input:
// 1 149,6088607594936;1 179,0666666666666
// Sample Output:
// 0.9750
func main() {
	var text1, text2 string

	text, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil && err != io.EOF {
		fmt.Println("Error EOF")
	} else {
		strNumberArray := strings.Split(text, ";")
		text1 = filterStr2(strNumberArray[0])
		text2 = filterStr2(strNumberArray[1])

		float1, _ := strconv.ParseFloat(strings.ReplaceAll(text1, ",", "."), 64)
		float2, _ := strconv.ParseFloat(strings.ReplaceAll(text2, ",", "."), 64)

		fmt.Printf("%.4f", float1/float2)
	}
}
