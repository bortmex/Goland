package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 3/8 Работа с датой и временем
// На стандартный ввод подается строковое представление даты и времени в следующем формате:
// 1986-04-16T05:20:00+06:00
// Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:
// Wed Apr 16 05:20:00 +0600 1986
// Для более эффективной работы рекомендуется ознакомиться с документацией о содержащихся в модуле time константах.
func main() {
	reader := bufio.NewReader(os.Stdin)
	all, _ := io.ReadAll(reader)
	parse, _ := time.Parse(time.RFC3339, strings.ReplaceAll(string(all), "\n", ""))
	fmt.Println(parse.Format(time.UnixDate))
}
