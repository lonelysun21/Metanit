package main

//3. Тип функции. Функция как параметр и результат функции
// У функции есть свой тип, например add(int,int) int.
// Функция может быть как параметр для другой функции
// Функция как результат другой функции

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {

	var f func(int, int) int = add
	fmt.Println(f(3, 4)) // 7

	var x = f(4, 5) // 9
	fmt.Println(x)
}

/*
   f := add        //или так var f func(int, int) int = add
   fmt.Println(f(3, 4))        // 7

   f = multiply    // теперь переменная f указывает на функцию multiply
   fmt.Println(f(3, 4))        // 12

*/
