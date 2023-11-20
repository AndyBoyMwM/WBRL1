package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.
func main() {
	var startArr = []int{2, 4, 6, 8, 10}

	wgSoloSquare(startArr)
	wgGroupSquare(startArr)
	chGroupSquare(startArr)
}

func wgSquare(n int, wg *sync.WaitGroup) {
	// уменьшаем счетчик в группе wg на 1, сигнал, что элемент закончил выполнение
	defer wg.Done()
	sq := n * n
	fmt.Println(sq)
}

func chSquare(n int, ch chan<- int, i int) {
	sq := n * n
	fmt.Println(sq)
	// пишем  в канал
	ch <- i
}

// с использованием WaitGroup, на каждый элемент массива
func wgSoloSquare(arr []int) {
	// создаем структуру wg
	var wg sync.WaitGroup
	fmt.Println("wgSoloSquare:")

	for _, n := range arr {
		// увеличиваем счетчик в группе wg на 1
		wg.Add(1)
		// запуск горутины на каждый элемент массива
		go func(n int) {
			// уменьшаем счетчик в группе wg на 1
			defer wg.Done()
			nn := n * n
			fmt.Println(nn)
		}(n)
	}
	// ждем, пока счетчик структуры wg не будет равен 0
	wg.Wait()
}

// с использованием WaitGroup, сразу на все элементы массива
func wgGroupSquare(arr []int) {

	wg := &sync.WaitGroup{} // создаем структуру wg
	wg.Add(len(arr))        // увеличиваем счетчик в группе wg на размер массива inarr

	fmt.Println("wgGroupSquare:")
	// когда отработает горутина она уменьшит счетчик wg на 1
	for _, numb := range arr {
		go wgSquare(numb, wg) // передаем массив по индексу и wg, после завершения горутира уменьшит wg на 1
	}
	// ждем, пока счетчик структуры wg не будет равен 0
	wg.Wait()
}

// с использованием каналов
func chGroupSquare(arr []int) {
	fmt.Println("chGroupSquare:")

	// создаем буферизированный канала ch
	ch := make(chan int, len(arr))

	for i, n := range arr {
		go chSquare(n, ch, i) // передаем значение массива, канал, индекс массива
	}

	// читаем каналы, для завершения работы горутин
	for i := 0; i < len(arr); i++ {
		<-ch
	}

	// закрываем канал
	close(ch)
}
