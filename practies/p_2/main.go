package main

// Метода

import (
	"fmt"
)

type library []string

func (l library) print() {
	for _, i := range l {
		fmt.Println(i)
	}
}

func main() {
	var lib library = library{"book1", "book2", "book3"}
	lib.print()
}
