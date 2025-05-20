package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

var s = myStruct{
	Name:   "John Connor",
	Age:    35,
	Status: true,
	Values: []int{15, 11, 37},
}

var myS myStruct

var jsons = []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

func main() {
	data, _ := json.Marshal(s)
	fmt.Printf("%s\n", data)

	data1, _ := json.MarshalIndent(s, "", "\t")
	fmt.Printf("%s\n", data1)

	json.Unmarshal(jsons, &myS)
	fmt.Printf("%v\n", myS)
}
