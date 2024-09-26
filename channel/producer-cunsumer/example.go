package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

// 차체 생산
func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports Car"
			tireCh <- car

		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

// 바퀴 설치
func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter Tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

// 도색
func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"

		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete\nCar : %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory\n")
}

/*
📌 채널로 생산자 소비자 패턴 구현하기
- 고루틴에서 뮤텍스를 사용하지 않는 방벙 중 -> 채널을 이용해 역할을 나누는 방법


📌 `컨베이터 벨트 시스템`
-  예를 들어 자동차 공장에서 자동차를
- `차체 생산` -> `바퀴 설치` -> `도색` -> `완성` 단계를 거쳐 생산한다고 가정하자.
- 각 공정에 1초가 걸린다고 보면 자동차 한 대를 만드는 데 3초가 걸린다.
- 그런데 3명이 공정 하나씩 처리하면 첫 차 생산에만 3초가걸리고 그 뒤론 1초마다 하나씩 생산할 수 있다.

- 작업자 간 자동차 전달은 컨테이어 벨트를 통해서만 이루어진다.
- 자동차는 데이터, 컨베이어 벨트는 채널로 볼 수 있다.


📌 생산자 소비자 패턴 (Producer Consumer Pattern)
- 한쪽에서 데이터를 생성해서 넣어주면 다른 쪽에서 생성된 데이터를 빼서 사용하는 방식을 생산자 소비자 패턴이ㅏㄹ고 한다.
- MakeBody() 루틴이 생산자
- InstallTire() 루틴은 소비자, PaintCar() 루틴에 대해서는 생산자
*/
