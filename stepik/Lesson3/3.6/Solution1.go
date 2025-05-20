package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// 9/9 JSON
// Данная задача ориентирована на реальную работу с данными в формате JSON.
// Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
// который был размещен на web-портале data.gov.ru.
// Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json, а сами данные,
// которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории на github.com.
// Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа записать
// сумму полей global_id всех элементов, закодированных в файле.

var myS1 []trb1maker

type trb1maker struct {
	Id int64 `json:"global_id"`
}

func main() {

	file, _ := os.Open("D:\\GoGoGo\\Projects\\Go\\myFirstProject\\stepik\\Lesson3\\3.6\\data-20190514T0100.json")
	data, _ := io.ReadAll(file)
	json.Unmarshal(data, &myS1)
	var count int64
	for _, glob_id := range myS1 {
		count += glob_id.Id
	}
	fmt.Println(count)
}
