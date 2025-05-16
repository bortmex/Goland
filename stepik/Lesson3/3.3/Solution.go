package main

import (
	"fmt"
	"strconv"
)

// 7/9 Анонимные функции
// Используем анонимные функции на практике.
// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
// которую внутри функции main в дальнейшем можно будет вызвать по имени fn.
// Функция на вход получает целое положительное число (uint), возвращает число
// того же типа в котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0,
// то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.
func main() {

	fn := func() func(uint33 uint) uint {
		return func(uint33 uint) uint {
			str := strconv.FormatUint(uint64(uint33), 10)
			strArray := []byte(str)
			var resultStr string
			for _, symbol := range strArray {
				if byte('2') == symbol || byte('4') == symbol || byte('6') == symbol || byte('8') == symbol {
					resultStr += string(symbol)
				}
			}
			if "" == resultStr {
				return 100
			} else {
				parseUint, _ := strconv.ParseUint(resultStr, 10, 64)
				return uint(parseUint)
			}
		}
	}()

	fmt.Println(fn(727178))
}
