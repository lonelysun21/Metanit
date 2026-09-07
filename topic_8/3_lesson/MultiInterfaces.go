package main

import (
	"fmt"
)

// 8.3 Multiple realization of interfaces.

// 1. Множественная реализация интерфейсов
// Одна структура может реализовать несколько интерфейсов. То есть тут тип Rectangle реализует методы обоих интерфейсов.

// 2. Наследование (Вложенные интерфейсы)
// В гоу нету наследования, но есть вложенные интерфейсы, то есть интерфейс в котором содержиться другие интерфейсы
// При этом во вложенном интерфейсе реализация также может осуществляться неявно
type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReaderWriter interface {
	Reader
	write(string) // неявно встраивается интерфейс Writer
}

type File struct {
	text string
}

func (f *File) read() {
	fmt.Println(f.text)
}
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func writeToStream(writer Writer, text string) {
	writer.write(text)
}

func main() {

	myFile := &File{}
	writeToStream(myFile, "Hello METANIT.COM")
	myFile.read()
}
