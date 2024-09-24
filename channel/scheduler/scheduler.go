package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

/*
📌 일정 간격으로 실행
- 메시지가 있으면 메시지를 빼와서 실행하고 그렇지 않다면 1초 간격으로 다른 일을 수행해야 한다는 가정을 해보자.
- time 패키지의 Tick() 함수로 원하는 시간 간격으로 신호를 보내주는 채널을 만들 수 있다.
*/
