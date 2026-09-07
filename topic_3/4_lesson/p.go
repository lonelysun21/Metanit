package main

// 4. Анонимные функции - функция без названия в главной функции main.
// Функция может быть как аргумент в другой функции (который мы позже создаем как функция или вписываем переменную которая хранит анонимную функцию)
// Анонимная функция как результат функции (возвращает функцию которую мы можем присвоить переменной и через переменную ее использовать)

import "fmt"

func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func main() {
	f := func(x, y int) int { return x + y }
	fmt.Println(f(5, 6))
	fmt.Println(f(7, 8))
	action(1, 2, func(x int, y int) int { return x + y })
}
