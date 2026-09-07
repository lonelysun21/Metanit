package main

import "fmt"

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {

	var tesla Car = Car{}
	var boing Aircraft = Aircraft{}
	tesla.move()
	boing.move()
}
