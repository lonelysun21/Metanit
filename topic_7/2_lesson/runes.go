package main

import "fmt"

// 7.1 Строки и руны.
// Руны - псевдоним для int32.

// Руны - это массив чисел, в котором числа это порядковый номер символа.
// Мы должны использовать %c - чтобы отформатировать числовое значение в символ
// string() - преобразование руны в строку.

func main() {
	//str1 := "Hello"
	str2 := "Привет"
	rune_slice := []rune(str2)
	for i := 0; i < len(rune_slice); i++ {
		fmt.Printf("%c ", rune_slice[i])
	}
	fmt.Print("\n")
	for _, value := range rune_slice {
		fmt.Printf("%c ", value)
	}
	fmt.Print("\n")
	fmt.Println(len(rune_slice))

	// Неизменяемость строки

	// В Go строки нельзя изменять напрямую, однако это можно сделать через руны

	fmt.Println(str2) // Обычная строка

	runes := []rune(str2) // превратили в руну

	runes[2] = 'ф' // поменяли символ

	str2 = string(runes) // вернули переменной str2

	fmt.Println(str2) // вывели новую строку
}
