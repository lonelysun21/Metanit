package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 11.6 Буферезированный ввод и вывод.

// 1. Запись через буфер
// В пакете bufio есть метода Writer:
/*
Write(): записывает срез байтов

WriteByte(): записывает один байт

WriteRune(): записывает один объект типа rune

WriteString(): записывает строку
*/

// Перед тем как использовать эти метода нужно даннные поместить в буфер через функцию Flush().

// 2. Для создания потока вывода через буфер применяется функция bufio.NewWriter():
//func NewWriter(w io.Writer) *Writer

// 3. Чтение данных через буфер
/*
Read(p []byte): считывает срез байтов и возвращает количество прочитанных байтов

ReadByte(): считывает один байт

ReadBytes(delim byte): считывает срез байтов из потока, пока не встретится байт delim

ReadLine(): считывает строку в виде среза байт

ReadRune(): считывает один объект типа rune

ReadSlice(delim byte): считывает срез байтов из потока, пока не встретится байт delim

ReadString(delim byte): считывает строку, пока не встретится байт delim
*/

func main() {
	file, err := os.Open("some.data")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Print(line)
	}
}
