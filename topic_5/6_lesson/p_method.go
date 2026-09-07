package main

// 5.5 Метода указателей
import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) update_age(new_age int) {
	p.age = new_age
}

func main() {
	tom := person{"Tom", 24}
	var p_tom *person = &tom
	fmt.Println("before", tom.age)
	p_tom.update_age(33)
	fmt.Println("after", tom.age)
}
