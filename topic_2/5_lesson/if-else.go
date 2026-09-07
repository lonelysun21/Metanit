package main

import "fmt"

func main() {
	a := 8
	b := 6

	if a > b { //if-else операция
		fmt.Println("a больше b")
	} else {
		fmt.Println("a меньше b")
	}

	switch a { // switch-case операция
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
		fallthrough // эта функция выполняет этот кейс не обращая внимание на условие
	default:
		fmt.Println("Значение может быть не точным")
	}
}
