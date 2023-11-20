package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения
// в канал, а с другой стороны канала — читать. По истечению N секунд
// программа должна завершаться.
func main() {
	var t int
	for {
		fmt.Print("Введите количество секунд: ")
		_, err := fmt.Scan(&t)
		if err != nil {
			fmt.Println("Ошибка ввода: введите целое число")
		} else {
			break
		}
	}

	// В канал delayWork пишем время завешшения в секундах, считанных в t
	delayWork := time.After(time.Duration(t) * time.Second)

	ch := make(chan int)

	// Запись в канал
	go func() {
		i := 0
		for {
			fmt.Printf("out: %v\n", i)
			ch <- i
			time.Sleep(200 * time.Millisecond)
			i++
		}
	}()

	// Чтение из канала
	go func() {
		for i := range ch {
			fmt.Printf("in : %v\n", i)
		}
	}()

	// Чтение из канала delayWork
	<-delayWork
}
