package main

// 5.3 Структуры как поля других структур
import (
	"fmt"
)

// Поля структур могут быть значением лругих структур.
// Таких имеется два вида: Вложенные структуры и встроенные структуры

// 1. Вложенные структуры - структура которая является частью, значением другой структуры

// 2. Встроенные структуры - определяются без имени поля, поэтому из поля прямым доступом продвигаются во внешнюю стркутуру

// 1. Пример вложенной структуры -->
type person struct {
	name string
	age  int
}

type account struct {
	login    string
	password string

	person_info person
}

func main() {
	tom := account{
		login:    "tom@hmail.com",
		password: "12345678",
		person_info: person{
			name: "Tom",
			age:  41,
		},
	}

	fmt.Println(tom) // {tom@hmail.com 12345678 {Tom 41}}
}

// 2. Пример встроенной структуры -->
/*
type person struct{  // вложенная структура person
    name string
    age int
}

type account struct{
    login string
    password string
    person
}

func main() {

    tom := account{
        "tom@hmail.com",
        "12345678",
        person {"Tom",  41,},
    }

    fmt.Println(tom)    // {tom@hmail.com 12345678 {Tom 41}}

    // обращение к полям встроенной структуры
    fmt.Println("Name: ", tom.person.name) // Name: Tom
}
*/
