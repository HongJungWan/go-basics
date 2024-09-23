package main

import "sync"

func main() {

}

func square(wg *sync.WaitGroup, ch chan int) {

}

/*
📌 버퍼를 가진 채널
- 내부에 데이터를 보관할 수 있는 메모리 영역을 버퍼라고 부른다.
- 버퍼가 다 차면, 버퍼가 없을 때와 마찬가지로 빈자리가 생길 때 까지 대기한다.
- 데이터를 제때 빼가지 않으면, 버퍼가 없을 때 처럼 고루틴이 멈추게 된다.

📌 고루틴에서 데이터를
*/
