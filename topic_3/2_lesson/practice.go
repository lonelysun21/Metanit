package main

//2. Возвращение результата из функции
import "fmt"

func main() {
	var age, name = add(4, 5, "Tom", "Simpson")
	fmt.Println(age)  // 9
	fmt.Println(name) // Tom Simpson
}

func add(x, y int, firstName, lastName string) (int, string) { // Возвращение нескольких значений
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

/*
func main() {
	var a = add(4, 5)  // 9
	var b = add(20, 6) // 26
	fmt.Println(a)
	fmt.Println(b)
}

func add(x, y int) int { // Возвращаем значение с типом int
	return x + y
}

func add1(x, y int) (z int) { // Именованные возвращаемые переменные (z возвращаем)
	z = x + y
	return
}

*/
