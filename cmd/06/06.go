package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.
func main() {

	//1 context.WithCancel - cancelFunc()
	ctxCancelFunc, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	go func() {
		fmt.Println("version 4 in")
		for {
			select {
			case <-ctxCancelFunc.Done():
				fmt.Println("version 4 out")
				return
			}
		}
	}()

	//2 context.WithTimeout - cancelTimeOut(), по времени, 5 секунд
	ctxTimeOut, cancelTimeOut := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeOut()

	go func() {
		fmt.Println("version 5 in")
		for {
			select {
			case <-ctxTimeOut.Done(): //если timeout "истек"
				fmt.Println("version 5 out")
				return
			}
		}
	}()

	//3 context.WithDeadline - cancelDeadline(), по времени дедлайна, 5 секунд
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancelDeadline()

	go func() {
		fmt.Println("version 7 in")
		for {
			select {
			case <-ctxDeadline.Done(): //если время превышено
				fmt.Println("version 7 out")
				return
			}
		}
	}()

	//4 с закрытием ch, из которого читаем в range
	ch := make(chan int)
	go func() {
		fmt.Println("version 6 in")
		for i := range ch {
			_ = i
		}
		fmt.Println("version 6 out")
	}()

	//5 по времени, 5 секунд
	timer := time.NewTimer(5 * time.Second)
	go func() {
		fmt.Println("version 1 in")
		for {
			select {
			case <-timer.C:
				fmt.Println("version 1 out")
				return
			}
		}
	}()

	//6 пока не появилась запись в off канале
	off := make(chan interface{})

	go func() {
		fmt.Println("version 2 in")
		for {
			select {
			case <-off:
				fmt.Println("version 2 out")
				return
			}
		}
	}()

	// 7 пока канал off не закрыт
	go func() {
		fmt.Println("version 3 in")
		for {
			select {
			case _, ok := <-off:
				if !ok {
					fmt.Println("version 3 out")
					return
				}
			}
		}
	}()

	time.Sleep(4 * time.Second)
	cancelFunc() // context.WithCancel - cancelFunc(), version 1
	close(ch)    // 4 close ch
	off <- true  // 6 пока не появилась запись в off канале
	close(off)   // 7 пока канал off не закрыт
	time.Sleep(4 * time.Second)
}
