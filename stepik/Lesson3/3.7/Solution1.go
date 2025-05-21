package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 4/8 Работа с датой и временем
// На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
// 2020-05-15 08:00:00
// Если время события до обеда (13-00), то ничего менять не нужно,
// достаточно вывести дату на стандартный вывод в том же формате.
// Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день,
// а затем вывести на стандартный вывод в том же формате.
// Sample Input:
// 2020-05-15 08:00:00
// Sample Output:
// 2020-05-15 08:00:00
func main() {
	reader := bufio.NewReader(os.Stdin)
	all, _ := io.ReadAll(reader)
	layout := "2006-01-02 15:04:05"
	parse, _ := time.Parse(layout, strings.ReplaceAll(string(all), "\n", ""))
	if parse.Hour() < 13 {
		fmt.Println(parse.Format(layout))
	} else {
		add := parse.Add(time.Hour * 24)
		fmt.Println(add.Format(layout))
	}
}
