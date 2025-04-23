package main

import "fmt"

type MyStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (myStr *MyStruct) Shoot() bool {
	if !myStr.On {
		return false
	}
	if myStr.Ammo != 0 {
		myStr.Ammo--
		return true
	} else {
		return false
	}
}

func (myStr *MyStruct) RideBike() bool {
	if !myStr.On {
		return false
	}
	if myStr.Power != 0 {
		myStr.Power--
		return true
	} else {
		return false
	}
}

// 8/8 Структуры
// В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции,
// как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.
// Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
// У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
// Если значение On == false, то оба метода вернут false.
// Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет,
// то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
func main() {
	test := new(MyStruct)
	fmt.Println(test)
}
