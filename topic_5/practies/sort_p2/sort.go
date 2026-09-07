package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Scan(&m)
	sort.Ints(slice)
	fmt.Println(slice)
	if sort.SearchInts(slice, m) == slices.Index(slice, m) {
		fmt.Println(sort.SearchInts(slice, m))
	} else {
		fmt.Println("Такое значение отсутсвует, но значение стояло бы на позиции:", sort.SearchInts(slice, m))
	}
}
