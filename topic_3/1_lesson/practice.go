package main

// 1. Функции и их параметры

import "fmt"

func main() {
	sum(5, 6)
	sum(7, 8, 9, 10, 11)
	sum(10, 9, 8, 7)
	sum([]int{1, 2, 3}...)
}
func sum(numbers ...int) { //Функция для сложения
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Sum:", sum)
}

/*
func sum(numbers ...int) { //Функция для сложения
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Sum:", sum)
}
*/
