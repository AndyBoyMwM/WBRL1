package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

1) var justString string - глобальная переменная, плюс не используется
2) руна может занимать не один байт, срезав слайс из 100 байт justString = v[:100], мы оставим часть строки "hugeStrng",
   которую сборщик мусора не сможет удалить, пока строка "hugeStrng" в области видимости.
*/

func main() {
	var justString0, justString1, justString2 string

	hugeStrng := createHugeString(1 << 10)

	// срез по количеству байт
	justString0 = hugeStrng[:100]
	fmt.Println("justString = v[:100] ", justString0)
	if reflect.DeepEqual([]byte(justString0), hugeStrng[:len(justString0)]) {
		fmt.Println("Да, из hugeStrng было взято ровно столько байт, сколько содержится в justString0")
	} else {
		fmt.Printf("Из hugeStrng не взято %v байт, в justString0 взято %v байт", len(hugeStrng), len(justString0))
	}
	fmt.Println()

	justString1 = strings.Clone(hugeStrng[:100]) // 1: strings.Clone
	fmt.Println("strings.Clone(v[:100]) ", justString1)
	if reflect.DeepEqual([]byte(justString1), hugeStrng[:len(justString1)]) {
		fmt.Println("Да, из hugeStrng было взято ровно столько байт, сколько содержится в justString1")
	} else {
		fmt.Printf("Из hugeStrng не взято %v байт, в justString1 взято %v байт", len(hugeStrng), len(justString1))
	}
	fmt.Println()

	justString2 = string([]rune(hugeStrng)[:100]) // 2: срезаем по количеству рун
	fmt.Println("string([]rune(v)[:100]) ", justString2)
	if reflect.DeepEqual([]byte(justString2), hugeStrng[:len(justString2)]) {
		fmt.Println("Да, из hugeStrng было взято ровно столько байт, сколько содержится в justString2")
	} else {
		fmt.Printf("Из hugeStrng не взято %v байт, в justString2 взято %v байт", len(hugeStrng), len(justString2))
	}

	fmt.Println()
}
func createHugeString(size int) string {
	return strings.Repeat("比", size)
}
