package main

import (
	"fmt"
)

func main() {
	var str = "Hello"
	var intval = 5
	var floatval = 5.9083
	var boolval = true
	fmt.Printf("Value: %v \n", intval) // вывод значений в формате синтаксиса гоу
	fmt.Printf("Value: %v \n", str)
	fmt.Printf("Value: %v \n", floatval)
	fmt.Printf("Value: %v \n", boolval)
	// Если я хочу форматировать и в середину записи вставлять значение, то можно просто использовать (%v)

	var b bool = 4 > 5 && 5 > 4 // and
	var c bool = 4 > 5 || 5 > 4 // or
	fmt.Printf("B: %v, C: %v \n", b, c)

	strh := "Hello"
	for index, value := range strh {
		fmt.Printf("Index: %v, Value: %c \n", index, value) // %c для вывода символов представленных числовым кодом
	}
}
