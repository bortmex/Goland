package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 8/13 Решение задач
// Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
// на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации.
// Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
// Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
// По данному числу определите его цифровой корень.
func main() {

	var result int
	var strResult string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	for true {
		resultArray := strings.Split(str, "")
		for i := 0; i < len(resultArray); i++ {
			resultInt, _ := strconv.Atoi(resultArray[i])
			result += resultInt
		}
		strResult = strconv.Itoa(result)
		if len(strResult) < 2 {
			break
		} else {
			str = strResult
			strResult = ""
			result = 0
		}
	}

	fmt.Println(strResult)
}
