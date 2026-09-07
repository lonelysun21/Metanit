package main

import (
	"fmt"
)

// 8.2 Реализация интерфейсов для указателей структур

// 1.
// Если же говорить проще то в тип данных. интерфейс может принимать не только структуру, но и указатель на структуру и все.
/*
file := File{"Hello METANIT.COM"}
    // так можно
    read_data(file)

    p_file := &file  // указатель на File
    // и так можно
    read_data(p_file)
*/

// 2.
// Однако если мы в методе укажем не структуру, а указатель на структуру, то интерфейс сможет принять только указатель на структуру, а при передаче структуре выдаст ошибку.
/*
 file := File{"Hello METANIT.COM"}
    p_file := &file  // указатель на File

    // указатель можно передать
    read_data(p_file)

    // а так нельзя
    // read_data(file)   // ! Ошибка
*/

// 3.
// Реализация интерфейсов для указателей может быть полезна если мы хотим изменить поле структуры, ведь
// Используя указатель мы напрямую обращаемся к полям структуры, даже через функции.

type Writer interface {
	write(string)
}

type File struct {
	text string
}

// реализация интерфейса Writer для *File
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func main() {

	myFile := File{"Undefined"}
	myFile.write("Hello World")
	fmt.Println(myFile.text) // Hello World
}
