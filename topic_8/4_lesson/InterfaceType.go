package main

import (
	"fmt"
)

// 8.4 Проверка типа интерфейса

// 1. Проверям тип одного интерфейса
// Мы можем проверить реализовывает ли структура интерфейс или нет при помощи такой конструкции
// t, ok := i.(T)
// value, ok := shape.(Rectangle) <-- пример

// интерфейс перемещения объекта
type Movable interface {
	move()
}

type Rectangle struct {
	x, // X-координата левого верхнего угла
	y, // Y-координата левого верхнего угла
	width, // ширина
	height int // высота
}

type Circle struct {
	x, // X-координата центра круга
	y, // Y-координата центра круга
	radius int // радиус круга
}

// реализация интерфейса Movable для Rectangle
func (r Rectangle) move() {
	fmt.Println("Перемещаем прямоугольник")
}

// реализация интерфейса Movable для Circle
func (c Circle) move() {
	fmt.Println("Перемещаем круг")
}

func main() {

	var shape Movable = Rectangle{x: 20, y: 10, width: 150, height: 100}
	//move_object(shape)

	// проверяем, реализует ли структура Rectangle интерфейс Movable
	value, ok := shape.(Rectangle)
	fmt.Println(ok)    // true
	fmt.Println(value) // {20 10 150 100}
}

// 2. Проверяем тип нескольких интерфейсов
// Данная проверка выполняется через конструкцию switch-case
/*
switch value := i.(type){

    case T1:        // Действия, если value представляет тип T1

    case T2:        // Действия, если value представляет тип T2

    .......................................................
    case TN:        // Действия, если value представляет тип TN

    default:    // если ни один из типов в case не соответствуют v
}
*/

/*
func check(i interface{}) {

    switch value := i.(type) {

        case Rectangle:
            fmt.Println("Type: Rectangle. Value: ", value)

        case Circle:
            fmt.Println("Type: Circle. Value: ", value)

        default:
            fmt.Println("Type: Undefined")
    }
}
*/
