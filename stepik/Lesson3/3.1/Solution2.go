package main

import "fmt"

// 7/7 Отображения (map)
func main() {
	var friends0fDima []string
	friends0fSemyon := make([]string, 3)
	friends0fDima = append(friends0fDima, "Костя", "Семён", "Таня")
	friends0fSemyon = append(friends0fSemyon, "Валера", "Таня", "Дима")
	friends := map[string][]string{
		"Dima":   friends0fDima,
		"Semyon": friends0fSemyon,
		"Костя":  nil,
	}
	_, value := friends["Костя"]
	fmt.Print(value, " ")
	delete(friends, "Dima")
	fmt.Print(friends0fSemyon[3])
}
