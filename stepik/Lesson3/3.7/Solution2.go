package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 7/8 Работа с датой и временем
// На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
// Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.
// Sample Input:
// 13.03.2018 14:00:15,12.03.2018 14:00:15
// Sample Output:
// 24h0m0s
func main() {
	reader := bufio.NewReader(os.Stdin)
	all, _ := io.ReadAll(reader)
	all1 := strings.ReplaceAll(string(all), "\n", "")
	arrayAll := strings.Split(all1, ",")
	layout := "02.01.2006 15:04:05"
	date1, _ := time.Parse(layout, arrayAll[0])
	date2, _ := time.Parse(layout, arrayAll[1])
	sub := date1.Sub(date2).Abs()
	fmt.Println(sub)
}
