package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
type mcCountr struct {
	mut sync.Mutex
	cnt int
}

func main() {
	var ac int64           // atomic
	mc := mcCountr{cnt: 0} // mutex
	var wg sync.WaitGroup
	wgNums := 500

	wg.Add(wgNums)
	for i := 0; i < wgNums; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				atomic.AddInt64(&ac, 1)
				mc.mut.Lock()
				mc.cnt++
				mc.mut.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("atomic", ac)
	fmt.Println("mutex ", mc.cnt)
}
