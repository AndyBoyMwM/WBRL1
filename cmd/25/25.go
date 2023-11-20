package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func main() {
	fmt.Println("begin sleeping...")
	sleep(2)
	fmt.Println("end sleeping")
}

func sleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}
