package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.
func main() {
	arr := []int{1, 2, 9, 4, 3, 3, 7, 8, 9}
	findNum := 7

	sort.Ints(arr)
	ind := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= findNum
	})
	fmt.Printf("индекс = %v, число = %v, arr = %v\n", ind, findNum, arr)
}
