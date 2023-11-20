package main

import (
	"fmt"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из
// канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
func main() {
	var cntRut int
	for {
		fmt.Print("Введите количество горутин: ")
		_, err := fmt.Scan(&cntRut)
		if err != nil {
			fmt.Println("Ошибка ввода: введите целое число")
		} else {
			break
		}
	}

	// создаем канал
	workerIn := make(chan int)

	// создаем горутины
	for i := 0; i < cntRut; i++ {
		go runWorker(i, workerIn)
	}

	// запись в канал
	for {
		workerIn <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func runWorker(workerNum int, chIn <-chan int) {
	for {
		fmt.Printf("goroutine #%v: = %vs\n", workerNum, <-chIn)
	}
}
