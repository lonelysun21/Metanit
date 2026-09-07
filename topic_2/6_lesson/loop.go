package main

//11. Циклы
// Рабочая программа таблицы умножения
import "fmt"

func main() {
	var arr [10][10]int

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			arr[i][j] = i * j
		}
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Println(i, "*", j, "=", arr[i][j], "\t")
		}
		fmt.Println()
	}
}
