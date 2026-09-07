package main

// 4.1 Указатели
import (
	"fmt"
)

func main() {
	// Массивы указателей
	var a, b, d int = 1, 2, 4
	var p_nums [4]*int // Массив указателей
	p_nums[0] = &a
	p_nums[1] = &b
	p_nums[3] = &d

	p_n := [4]*int{&a, &b, nil, &d}

	fmt.Println(p_nums, p_n) // [0xc000010100 0xc000010108 <nil> 0xc000010110]
}

/*
func main() {
	// Указатели на укахатели
	var a int = 9

	var p_a *int = &a // p_a - указатель на переменную a

	// обращаемся к указателю p_a
	fmt.Println("Value of a:", *p_a)  // Value of a: 9
	fmt.Println("Address of a:", p_a) // Address of a: 0xc000010100

	var p_p_a **int = &p_a // p_p_a - указатель на переменную p_a

	// обращаемся к указателю p_a
	fmt.Println("Value of p_a:", *p_p_a)  // Value of p_a: 0xc000010100
	fmt.Println("Address of p_a:", p_p_a) // Address of p_a: 0xc000076040
}
*/
