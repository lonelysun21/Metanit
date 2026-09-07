package main

//10. Массивы.
import "fmt"

func main() {
	var numbers1 [5]int = [5]int{1, 2, 3, 4, 5}            //обычный массив
	numbers2 := [...]int{6, 7, 8, 9, 10}                   // динамический массив
	colors := [...]string{3: "blue", 0: "red", 1: "green"} //массив с выбранными индексами
	numbers3 := [4][2]int{{1, 2}, {3, 4}, {5, 6}}          //двумерный массив
	num1 := numbers1                                       //Копирование одного массива в другой массив
	num1[1] = 11
	fmt.Println(numbers1)
	fmt.Println(numbers2)
	fmt.Println(colors[3])
	fmt.Println(numbers3)
	fmt.Println(len(numbers1)) //длина первого массива
	fmt.Println(num1)
}
