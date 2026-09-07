package main

//11. Циклы
import "fmt"

func main() {

	str := "Hello"

	for index, value := range str {
		fmt.Printf("Index: %d, Value: %c\n", index, value) // Перебираем список значения стринг по буквам
	}

	for _, value := range str {
		fmt.Printf("%c", value)
	}
}

// break OuterLoop - Выходит из внешнего цикла и заврешает его
