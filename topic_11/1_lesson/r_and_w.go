package main

// 11.1 Операции ввода и вывода. Reader and Writer

// 1. Поток данных
// Поток данных в гоу это байтовый срез, из которых можно считывать байты или заносить данные.
// Кючевые типами для работы с потоками является интерфейсы Reader and Writer из библиотеки io.

type Reader interface {
	Read(p []byte) (n int, err error)
} // Метод read возвращает общее количество считанных байтов и информацию об ошибке если она возникает. Если в потоке нету данных, то возвращает io.EOF

// 2. io.Reader
/*
package main
import (
"fmt"
"io"
)

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error){
    count := 0
    for i := 0; i < len(ph); i++{
        if(ph[i] >= '0' && ph[i] <= '9'){
            p[count] = ph[i]
            count++
        }
    }
    return count, io.EOF
}

func main() {
    phone1 := phoneReader("+1(234)567 9010")
    phone2 := phoneReader("+2-345-678-12-35")

    buffer := make([]byte, len(phone1))
    phone1.Read(buffer)
    fmt.Println(string(buffer))     // 12345679010

    buffer = make([]byte, len(phone2))
    phone2.Read(buffer)
    fmt.Println(string(buffer))     // 23456781235
}
*/

// 3. io.Writer
type Writer interface {
	Write(p []byte) (n int, err error)
} // Предназначен для записи в поток.
/*
package main
import "fmt"

type phoneWriter struct{ }

	func (p phoneWriter) Write(bs []byte) (int, error){
	    if len(bs) == 0 {
	         return 0, nil
	   }
	   for i := 0; i < len(bs); i++{
	        if(bs[i] >= '0' && bs[i] <= '9'){
	            fmt.Print(string(bs[i]))
	        }
	    }
	    fmt.Println()
	    return len(bs), nil
	}

func main() {

	    bytes1 := []byte("+1(234)567 9010")
	    bytes2 := []byte("+2-345-678-12-35")

	    writer := phoneWriter{}
	    writer.Write(bytes1)
	    writer.Write(bytes2)
	}
*/
func main() {

}
