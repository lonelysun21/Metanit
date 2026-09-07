package main

// 5.2 Структуры
import (
	"fmt"
)

type person struct { // структура которая присваиваеться переменной (как я понял: person, это тип данных или кратко псевдоним, который имеет тип данных struct)
	name string
	age  uint
}

var aru struct { // Анонимная структура
	name string
	age  uint
}

type sun struct { // Анонимные поля структур
	string // В анонимных полях можно указать только один тип данных
	uint   // Я не могу указать два раза подряд string
} // Однако для второго стринг мы можем просто указать имя переменной и все

func main() {
	// Структура
	var Sanzhar person = person{"Oleg", 20}
	fmt.Println(Sanzhar.name, Sanzhar.age)

	// Анонимные структура
	aru.name = "Aru"
	aru.age = 16
	fmt.Println(aru.name, aru.age)

	// Анонимные поля структур
	san := sun{"San", 20}
	fmt.Println(san)
}
