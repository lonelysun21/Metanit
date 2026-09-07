package main

// 5. Замыкание
// Замыкание это когда внутри функции параметр функции или его переменные запоминают значение, а внутреняя возвращаемая функция может работать с ней

import "fmt"

func multiply(n int) func(int) int {

	return func(m int) int { return n * m }
}

func main() {

	fn := multiply(5)
	result1 := fn(6)     // 30
	fmt.Println(result1) // 30

	result2 := fn(5)     // 25
	fmt.Println(result2) // 25
}
