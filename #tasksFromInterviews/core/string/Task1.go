package main

import "fmt"

func main() {
	str := "это строка"
	bs := []byte(str[7:])

	fmt.Printf("Так срез выглядит: %s\n", bs)
	fmt.Printf("Так байтовый срез выглядит внутри: %v\n", bs)

	for i := range bs {
		if bs[i]%2 == 0 {
			bs[i] = bs[i] + 1
		} else {
			bs[i] = bs[i] - 1
		}
	}

	fmt.Printf("Исходный массив строки: %s\n", str)
	fmt.Printf("Изменённый байтовый срез в виде строки: %s", bs)
}
