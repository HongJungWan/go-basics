package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)

}

/*
- 모든 프로그램은 고루틴을 최소한 하나는 가지고 있다.
- 바로 메인 루틴이다. 이 고루틴은 main() 함수와 함께 시작되고 main() 함수가 종료되면 종료된다. 또 메인 루틴이 종료되면 프로그램 또한 종료하게 된다.
*/

/*
📌 고루틴 동작 원리
- ㄴ
*/
