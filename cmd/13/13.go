package main

import (
	"fmt"
	"sync/atomic"
)

// Поменять местами два числа без создания временной переменной.
func main() {
	var x, y int64 = 100, 200
	fmt.Printf("start values: x = %v, y = %v\n", x, y)

	y = atomic.SwapInt64(&x, y)
	fmt.Printf("start -> atomic: x = %v, y = %v\n", x, y)

	x, y = y, x
	fmt.Printf("atomic -> swap: x = %v, y = %v\n", x, y)

	x ^= y
	y ^= x
	x ^= y
	fmt.Printf("swap -> bit   : x = %v, y = %v\n", x, y)
}
