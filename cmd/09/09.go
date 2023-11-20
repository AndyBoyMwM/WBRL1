package main

import (
	"fmt"
)

// 9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("in: ", arr)
	ch := genArr(arr)
	out := sqr(ch)

	fmt.Printf("out: ")
	for val := range out {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func genArr(arr []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range arr {
			out <- i
		}
		close(out)
	}()
	return out
}

func sqr(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}
