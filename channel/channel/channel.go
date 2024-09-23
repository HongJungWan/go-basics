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
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // 먼저 채널에서 데이터를 빼오려고 시도, 채널에 아무런 데이터가 없으면 데이터가 들어올 때 까지 대기한다.

	time.Sleep(time.Second)
	fmt.Println("Square: %d\n", n*n)
	wg.Done()
}

/*
📌 채널과 컨텍스트는 Go 언어에서 동시성 프로그래밍을 도와주는 기능이다.
- 채널은 고루틴 간 메시지를 전달하는 메시지 큐다. 채널을 사용하면 뮤텍스 없이 동시성 프로그래밍이 가능하다.
- 메시지 큐에 메시지들은 차례대로 쌓이게 되고 메시지를 읽을 때는 맨 처음 온 메시지부터 차례대로 읽게 된다.

- 컨텍스트는 고루틴에 작업을 요청할 때 작업 취소나 작업 시간 등을 설정할 수 있는 작업 명세서 역할을 한다.

- 채널과 컨텍스트를 사용해 특정 데이터를 전달하거나 특정 시간 동안만 작업을 요청하거나 작업 도중에 작업 취소를 요청할 수 있다.
*/
