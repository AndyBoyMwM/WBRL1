package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.
func main() {
	var nRut int

	for {
		fmt.Print("Введите количество горутин: ")
		_, err := fmt.Scan(&nRut)
		if err != nil {
			fmt.Println("Ошибка ввода, введите целое число")
		} else {
			break
		}
	}

	mutexMap(nRut)
	syncMap(nRut)
}

func mutexMap(nRut int) {
	intMap := make(map[int]int)
	mtx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("mutex:")

	for i := 0; i <= nRut; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mtx.Lock()
			intMap[i] = rand.Intn(99)
			mtx.Unlock()
		}(i)
	}
	wg.Wait()

	for i := range intMap {
		fmt.Printf("key: %v, value: %v\n", i, intMap[i])
	}
}

func syncMap(nRut int) {
	intMap := &sync.Map{}
	wg := &sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("sync:")

	for i := 0; i <= nRut; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			intMap.Store(i, rand.Intn(99))
		}(i)
	}
	wg.Wait()

	for i := 0; i <= nRut; i++ {
		val, _ := intMap.Load(i)
		fmt.Printf("key: %v, value: %v\n", i, val)
	}
}
