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
	close(ch)
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

/*
📌 버퍼를 가진 채널
- 내부에 데이터를 보관할 수 있는 메모리 영역을 버퍼라고 부른다.
- 버퍼가 다 차면, 버퍼가 없을 때와 마찬가지로 빈자리가 생길 때 까지 대기한다.
- 데이터를 제때 빼가지 않으면, 버퍼가 없을 때 처럼 고루틴이 멈추게 된다.

📌 채널을 다 사용하면 close(ch)를 호출해 채널을 닫고 채널이 닫혔음을 알려줘야 한다.
- 채널에서 데이터를 모두 빼낸 상태이고 채널이 닫혔으면 for range문을 빠져나가게 된다.
- 채널을 제때 닫아주지 않아서 고루틴에서 데이터를 기다리며 무한 대기하는 경우를 좀비 루틴 또는 고루틴 릭(leak)이라고 한다.
- 아무리 경량 스레드라고 해도 고루틴 또한 메모리와 성능을 차지하기 때문에 이런 좀비 루틴이 많아지면 메모리가 부족할 수 있다.
- 그래서 놀고 있는 고루틴이 없는지 잘 체크해봐야한다.
*/
