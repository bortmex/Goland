package main

import (
	"container/list"
	"fmt"
)

//	5/6 Списки
// Задача - перевот списка (reverse)
// Вам нужно реализовать функцию, которая принимает list и переворачивает порядок его элементов,
// так чтобы последний элемент стал первым, предпоследний — вторым, и так далее.

// ReverseList - функция для реверса списка
func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		reversedList.PushFront(elem.Value)
	}
	return reversedList
}

func printList(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
	fmt.Println()
}

func main() {

	task := list.New()

	for i := 0; i < 10; i++ {
		task.PushBack(i)
	}
	printList(task)
	// 0123456789

	printList(ReverseList(task))
	// 9876543210
}
