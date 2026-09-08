package main

import (
	"fmt"
	//"io"
	"os"
)

// 11.3 Стандартные потоки ввода и вывода os.Stdin, os.Stdout, os.Stderr, стандартный поток ввода, вывода и вывода ошибок.
// То есть благодаря потоку вывода мы можем вывести весь скопированный текст из файла на вывод в консоли.
func main() {
	file, err := os.Create("hello2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprint(file, "Hello Aruzhan!") // данная функция выводит текст в указанный нами адрес вначале (в файл, в сетевой поток и.т.д)
	//io.Copy(os.Stdout, file)
	fmt.Fprintln(os.Stdout, "Hello There!")
}
