package main

import (
	"fmt"
	"reflect"
)

// 5.7 Карты (maps)

// Карта или map (хэш-таблица или словарь) - структура данных "ключ - значение"
func main() {
	var people map[string]int
	fmt.Println(people) // map[]
	people1 := map[string]int{
		"tom":  1,
		"sara": 2,
		"bob":  3,
		"sam":  4,
	}
	fmt.Println(people1) // map[bob:3 sam:4 sara:2 tom:1]

	// 1. Обращение к элементам карты
	fmt.Println(people1["tom"])          // если 100% уверены что значение по такому ключу есть (если нету, то выведется ошибка)
	if value, ok := people1["bob"]; ok { // провяреям есть ли, если есть то выводим, без ошибки
		fmt.Println(value)
	}

	// 2. Создание карты через make()
	var people2 = make(map[string]int)
	// 3. Добавление и удаление элементов
	people2["Kate"] = 128 // добавляем новое значение в словарь
	delete(people2, "Kate")
	fmt.Println(people2)

	// 3. Сравнение карт
	// нужно подключить пакет "reflect"
	// возвращает true, если обе карты имеют одинаковые ключи и одинаковые связанные значения ключей
	// в противном случае возвращает false
	// функцию DeepEqual
	fmt.Println(reflect.DeepEqual(people1, people2))

}
