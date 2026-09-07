package main

import (
	"fmt"
	"os"
)

// 11.2 Создание и открытие файлов
// Чтобы работать с файлами можно использовать функциональность пакета os. Мы уже озгнакомились с интерфейсами io.Reader и io.Writer, которые туда входят

// 1.1 os.Create()
// Можно создать файл по указанному пути: file, err := os.Create("hello.txt")

// 1.2 os.Open()
// Ранее созданный файл можно открыть с помощью функции os.Open().  file, err := os.Open("hello.txt")

// 1.3 os.OpenFile()
// Функция открывает файл, а если его нету, то создает его. Функция принримает три. параметра:
// 1. путь к файлу, 2. режим открытия файла (для записи, для чтения), 3. разрешение для доступа к файлу.

// 1.4 Close().
// После работы с файлом, его нужно закрыть функцией Close().

// 1.5 os.Exit()
// Позволяет выйти из программы.

// 1.6 Name()
// Определен для типа os.File, и он возвращает название файла.
func main() {
	file, err := os.Create("hello.txt") // создаем файл
	if err != nil {                     // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	defer file.Close()       // закрываем файл
	fmt.Println(file.Name()) // hello.txt
}

// 2. Запись в файл
// WriteString() - чтобы заносить строку в файл

/*
package main
import (
    "fmt"
    "os"
)

func main() {
    text := "Hello Gold!"
    file, err := os.Create("hello.txt")

    if err != nil{
        fmt.Println("Unable to create file:", err)
        os.Exit(1)
    }
    defer file.Close()
    file.WriteString(text)

    fmt.Println("Done.")
}
*/

/*
package main
import (
    "fmt"
    "os"
)

func main() {
    data := []byte("Hello Bold!")
    file, err := os.Create("hello.bin")
    if err != nil{
        fmt.Println("Unable to create file:", err)
        os.Exit(1)
    }
    defer file.Close()
    file.Write(data)

    fmt.Println("Done.")
}
*/

// Чтение файла
// Позволяет получить содержимое файла в виде набора байтов

/*
package main
import (
    "fmt"
    "os"
    "io"
)

func main() {
    file, err := os.Open("hello.txt")
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    data := make([]byte, 64)

    for{
        n, err := file.Read(data)
        if err == io.EOF{   // если конец файла
            break           // выходим из цикла
        }
        fmt.Print(string(data[:n]))
    }
}
*/

// Пример из гоу рефреш
/*
fileIn, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error with opening first file")
		os.Exit(1)
	}
	defer fileIn.Close()

	fileOut, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error with opening second file")
		os.Exit(1)
	}
	defer fileOut.Close()

	data := make([]byte, 64)

	for {
		n, err := fileIn.Read(data)
		if err == io.EOF {
			break
		}
		fileOut.Write(data[:n])
	}

*/
