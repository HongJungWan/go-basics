package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

/*
- select문은 위와 같이 여러 채널을 동시에 기달리 수 있다. 만약 어떤 채널이라도 하나의 채널에서 데이터를 읽어오면
- 해당 구문을 실행하고 select문이 종료된다. 하나의 case만 처리되면 종료되기 때문에 반복해서 데이터를 처리하고 싶으면 for문과 함께 사용해야 한다.
*/
