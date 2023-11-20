package main

import "fmt"

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
// с использованием конкурентных вычислений.
func main() {
	startArr := []int{2, 4, 6, 8, 10}
	sqSum := getSqSum(startArr...)
	fmt.Println(sqSum)
}

func getSqSum(arr ...int) int {
	ch := make(chan int)
	go func() {
		for _, n := range arr {
			ch <- n * n
		}
		close(ch)
	}()

	sqSum := 0
	for sq := range ch {
		sqSum += sq
	}

	return sqSum
}
