package main

import "fmt"

func work(x int) int {
	return x
}

// 5/7 Отображения (map)
// Внутри функции main (объявлять функцию не нужно) необходимо написать программу:
// На стандартный ввод подается 10 целых чисел, разделенных пробелами (числа могут повторяться).
// Для чтения из стандартного ввода импортирован пакет fmt.
// Вам необходимо вычислить результат выполнения функции work для каждого из полученных чисел. Функция work имеет следующий вид:
// func work(x int) int
// Результаты вычислений, разделенные пробелами, должны быть напечатаны в строку.
// Однако работа функции work занимает слишком много времени.
// Выделенного вам времени выполнения не хватит на последовательную обработку каждого числа,
// поэтому необходимо реализовать кэширование уже готовых результатов и использовать их в работе.
// После завершения работы программы результат выполнения будет дополнен информацией о соблюдении установленного лимита времени выполнения.
func main() {
	var numbers [10]int
	cachi := make(map[int]int)

	for i := 0; i < 10; i++ {
		fmt.Scan(&numbers[i])
	}

	var results [10]int

	for i := 0; i < len(numbers); i++ {
		result, ok := cachi[numbers[i]]
		if ok {
			results[i] = result
		} else {
			result = work(numbers[i])
			results[i] = result
			cachi[numbers[i]] = result
		}
	}

	for result := range results {
		fmt.Print(results[result], " ")
	}

}
