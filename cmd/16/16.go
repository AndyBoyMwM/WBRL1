package main

import (
	"fmt"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	arr := []int{5, 5, 3, 2, 9, 1}
	fmt.Println("quickSort algo:")
	fmt.Println("in: ", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("out: %v\n\n", arr)
}

// Быстрая сортировка, Чарльз А. Р. Хоар
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	pivot := arr[(left+right)/2]

	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for pivot < arr[j] {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i, j = i+1, j-1
		}
	}
	quickSort(arr, left, j)
	quickSort(arr, i, right)
}
