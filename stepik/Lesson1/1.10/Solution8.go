package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 10/10 Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
// Входные данные
// Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
// Выходные данные
// Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
// Цифры выводятся в порядке их нахождения в первом числе.
func main() {
	var number1 int
	var number2 int
	result := []string{}

	fmt.Scan(&number1)
	fmt.Scan(&number2)

	str1 := strconv.Itoa(number1)
	str2 := strconv.Itoa(number2)

	split1 := strings.Split(str1, "")

	for i := 0; i < len(split1); i++ {
		if strings.ContainsAny(str2, split1[i]) {
			result = append(result, split1[i])
		}
	}

	justString := strings.Join(result, " ")
	fmt.Println(justString)
}
