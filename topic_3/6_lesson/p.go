package main

//6. Рекурсивные функции
import "fmt"

func sum(n uint) uint {

	if n == 1 {
		return n
	}
	return n + sum(n-1)
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibbonachi(n uint) uint {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibbonachi(n-1) + fibbonachi(n-2)
}

func main() {
	fmt.Println(sum(5))
	fmt.Println(factorial(5))
	fmt.Println(fibbonachi(6))
}

/*
func factorial(n uint) uint { // Функция факторияла

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {

	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(6)) // 720
}
*/

/*
func sum(n uint) uint { // Рекурсивные функции (сумма чисел)

	if n == 1 {
		return n
	}
	return n + sum(n-1)
}

func main() {

	fmt.Println(sum(4)) // 10
	fmt.Println(sum(5)) // 15
	fmt.Println(sum(6)) // 21
}

*/
