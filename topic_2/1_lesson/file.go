package main /* код хранится в файлах, а файлы должны пренадлежать пакетам*/
import "fmt" // подключаем пакет чтобы выводить данные
// 2.1 Файл, офомрление и объявление переменных
var sun1 int // как можно объвлять переменные
var sun2 int = 21
var sun3 = 22

var (
	hello string = "hello"
	age   int    = 20
) // объявление нескольких переменных в var

var (
	a int       = 28
	b float32   = 4.5
	c complex64 = 1 + 2i
	d bool      = true
	e string    = "Lonelysun \t ryichi"
) //различные типы данных

func main() { //все в программе выполняется в функции main

	sun4 := 23
	fmt.Println(sun4)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("%T", a)
}

/*Практика второй главы по этим темам:
Структура программы
Переменные
Типы данных
*/
