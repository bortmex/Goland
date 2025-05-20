package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

// 6/9 JSON
// На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
//
//	{
//	   "ID":134,
//	   "Number":"ИЛМ-1274",
//	   "Year":2,
//	   "Students":[
//	       {
//	           "LastName":"Вещий",
//	           "FirstName":"Лифон",
//	           "MiddleName":"Вениаминович",
//	           "Birthday":"4апреля1970года",
//	           "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//	           "Phone":"+7(948)709-47-24",
//	           "Rating":[1,2,3]
//	       },
//	       {
//	           // ...
//	       }
//	   ]
//	}
//
// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
// Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
// Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
//
//	{
//	    "Average": 14.1
//	}
//
// Как вы понимаете, для декодирования используется функция Unmarshal,
// а для кодирования MarshalIndent (префикс - пустая строка, отступ - 4 пробела).
var jsonData = []byte(`{
	"ID":134,
	"Number":"ИЛМ-1274",
	"Year":2,
	"Students":[
				{
				"LastName":"Вещий",
				"FirstName":"Лифон",
				"MiddleName":"Вениаминович",
				"Birthday":"4апреля1970года",
				"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
				"Phone":"+7(948)709-47-24",
				"Rating":[1,2]
				},
				{
				"LastName":"a",
				"FirstName":"b",
				"MiddleName":"w",
				"Birthday":"r",
				"Address":"h",
				"Phone":"p",
				"Rating":[1,2]
				}
			]
}`)

var myS myStruct

type myStruct struct {
	Students []students
}
type students struct {
	Rating []int
}
type answer struct {
	Average float64
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	all, _ := io.ReadAll(reader)
	json.Unmarshal(all, &myS)
	var countStudents int
	var countRate int
	for _, rates := range myS.Students {
		countStudents++
		countRate += len(rates.Rating)
	}
	resultFloat := float64(countRate) / float64(countStudents)

	resultFloatFromJson := answer{
		Average: resultFloat,
	}
	data1, _ := json.MarshalIndent(resultFloatFromJson, "", "    ")
	os.Stdout.Write(data1)
}
