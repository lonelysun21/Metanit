package main

import "fmt"

//6. Арифметические операции
//7. Условные выражения
//8. Поразрядные операции

func main() {
	bob := Cats{"bob", 7, 0.87}
	fmt.Println("Bob function is", bob.test())
}

type Cats struct {
	name      string
	age       int
	happiness float64
}

func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happiness
}

//Структуры
/**/

//Указатели
/*func main() {
	var x = 0
	pointer(&x)
	fmt.Println(x)
}

func pointer(x *int) {
	*x = 2
}*/

//Откладывание
/*func main() {
	defer two()
	one()
}

func one() {
	fmt.Println("1")
}

func two() {
	fmt.Println("2")
}*/

//Замыкание (те же функции, но создаются внутри других функций, а не отдельно)
/*var num int = 3
multiple := func() int {
	num *= 2
	return num
}
fmt.Println(multiple)*/

//Функции
/*func main() {
	var a = 29
	var b = 1
	a, b = summ(a, b)
	fmt.Println(a, b)
}

func summ(num_1 int, num_2 int) (int, int) {
	var res int
	res = num_1 + num_2
	var res1 = num_1 * num_2
	return res, res1
}*/

//Карты
/*webSites := make(map[string]float64)

webSites["Lonely_sun"] = 0.8
webSites["yandex"] = 0.99
fmt.Println(webSites["Lonely_sun"])*/

// Массивы
/*var arr [3]int
arr[0] = 45
arr[1] = 97
arr[2] = 76
fmt.Println(arr[1])

nums := [3]float64{4.23, 5.23, 6.23}
for i, value := range nums {
	fmt.Println(value, i)
}*/

// Циклы
/*var i = 1
for i <= 10 {
	fmt.Println(i)
	i++
}
*/

/*for i := 0; i <= 5; i++ {
	fmt.Println(i)
}*/

//Свитч-кейс
/*var age = 10
switch age {
case 5:
	fmt.Println("5")
case 15:
	fmt.Println("15")
case 10:
	fmt.Println("10")
default:
	fmt.Println("Вам неизвестно сколько вам лет")
}*/
