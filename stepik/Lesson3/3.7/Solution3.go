package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// вам это понадобится
const now = 1589570165

// 8/8 Работа с датой и временем
// На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).
// Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
// Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.
// Sample Input:
// 12 мин. 13 сек.
// Sample Output:
// Fri May 15 19:28:18 UTC 2020
func main() {
	reader := bufio.NewReader(os.Stdin)
	all, _ := io.ReadAll(reader)
	all1 := strings.ReplaceAll(string(all), "\n", "")
	minutes := strings.Split(all1, " мин. ")
	secundes := strings.Split(minutes[1], " сек.")
	//fmt.Scanf("%d мин. %d", &min, &sec)
	minutesGo, _ := strconv.Atoi(minutes[0])
	secundesGo, _ := strconv.Atoi(secundes[0])
	duration := time.Unix(now, 0).Add(time.Minute * time.Duration(minutesGo)).Add(time.Second * time.Duration(secundesGo))
	fmt.Println(duration.Format(time.UnixDate))
}
