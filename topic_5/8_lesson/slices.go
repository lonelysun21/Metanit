package main

import "fmt"

// 5.6 Срезы

func main() {
	// 1. Создание среза с помощью функции make
	//имя_среза := make([] тип_элементов_среза, длина_среза, емкость_среза)
	slice := make([]int, 5, 10)
	fmt.Println(slice)

	// 2. Двухмерные срезы
	slice_2d := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{5, 6},
	}
	fmt.Println(slice_2d)

	// 3. Добавление в срез
	users1 := []string{"Tom", "Alice", "Kate"}
	users1 = append(users1, "Bob") // append(slice, value)

	fmt.Println(users1) // [Tom Alice Kate Bob]

	// 4. Удаление элемента
	users2 := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	var n = 3
	users2 = append(users2[:n], users2[n+1:]...) // удаляем через append, как бы добавляя весь срез заново, пропуская не нужный элемент.
	fmt.Println(users2)

	// 5. Копирование среза
	// копирование происходит через функцию copy()

	/*
			Из-за начальной длины пустого списка при копирование скопирует 0 элементов
			slice1 := [] int {1, 2, 3, 4, 5, 6}
		    slice2 := [] int {}
		    copy(slice2, slice1)  // копируем из slice2 в slice2
		    fmt.Println(slice2)  // [] - slice2 пуст
	*/

	/*
		slice1 := [] int {1, 2, 3, 4, 5, 6}
		slice2 := make([]int, 3)
		copy(slice2, slice1)  // копируем из slice2 в slice2
		fmt.Println(slice2)  // [1 2 3]
	*/

	/*
		slice1 := [] int {1, 2, 3, 4, 5, 6}
		slice2 := make([]int, 8)
		copy(slice2, slice1)  // копируем из slice2 в slice2
		fmt.Println(slice2)  // [1 2 3 4 5 6 0 0]
	*/

	// 6. Сортировка среза
	// sort.Ints(), sort.Float64s() и sort.Strings()
	//"sort"  подключаем пакет sort
	// Example - sort.Ints(numbers) -- number - это срез чисел
	// Элементы сортируются в порядке возрастания

	// 7. Поиск в срезе
	// sort.SearchInts(), sort.SearchFloat64s() и sort.SearchStrings()
	//"sort"  подключаем пакет sort
	// поиск возможен если срез отсортирован!
	// Если элемент найден, то возвращает его индекс, если нет, то возвращает индекс, по которому он должен стоят в отсортированном списке.

	// 8. Срез в обратном порядке
	// Но перед переворачиванием среза, его надо обернуть в тип IntSlice/Float64Slice/StringSlice
	// с помощью соответственно функций sort.IntSlice/sort.Float64Slice/sort.StringSlice:

	/*
			intSlice := [] int {11, 22, 33, 44, 55, 66}
		    sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
		    fmt.Println(intSlice)  // [66 55 44 33 22 11]
	*/

	// 9. Сравнение срезов
	// Срезы считаются равными если у них один тип данных, и элементы все одинаковые
	// функцию DeepEqual()
	// нужен пакет reflect

	/*
		fmt.Println("slice1 == slice2:", reflect.DeepEqual(slice1, slice2))     // false
	    fmt.Println("slice1 == slice3:", reflect.DeepEqual(slice1, slice3))     // false
	    fmt.Println("slice1 == slice4:", reflect. DeepEqual(slice1, slice4))    // true
	*/

}
