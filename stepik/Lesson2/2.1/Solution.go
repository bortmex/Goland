package main

func vote(x int, y int, z int) int {

	var mapVote = map[int]int{}
	addMapAndCalculateSize(mapVote, x)
	addMapAndCalculateSize(mapVote, y)
	addMapAndCalculateSize(mapVote, z)

	if mapVote[0] > 1 {
		return 0
	} else {
		return 1
	}
}

func addMapAndCalculateSize(map1 map[int]int, value int) {
	valueInMap, ok1 := map1[value]
	if ok1 {
		map1[value] = valueInMap + 1
	} else {
		map1[value] = 1
	}
}

// 8/12 Функции
// Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.
// Входные данные
// В функцию передаются 3 числа - x, y и z (x, y и z равны 0 или 1).
// Выходные данные
// Необходимо вернуть значение функции от x, y и z.
func main() {
}
