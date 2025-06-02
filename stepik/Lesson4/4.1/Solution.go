package main

import (
	"container/list"
	"fmt"
)

func work(x int) int {
	return x
}

//	4/6 Списки
// Задача FIFO
// Реализуйте функции для очереди FIFO (First In, First Out (ПЕРВЫЙ пришел, ПЕРВЫЙ вышел)) с помощью списков.
// Должны быть данные функции:
// Push (добавление элемента)
// Pop (удаление элемента и его возврат)
// printQueue (печать очереди в одну строку без пробелов)

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	front := queue.Front()
	queue.Remove(front)
	return front
}

func printQueue(queue *list.List) {
	for elem := queue.Front(); elem != nil; elem = elem.Next() {
		fmt.Print(elem.Value)
	}
}

func main() {
}
