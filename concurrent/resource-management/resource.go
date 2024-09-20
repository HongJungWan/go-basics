package main

import (
	"fmt"
	"sync"
	"time"
)

// 고루틴이 같은 자원에 접근하지 않게 - 영역을 나누는 방법
type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}

	wg.Wait()
}

/*
📌 또 다른 자원 관리 기법
- 모든 문제는 같은 자원을 여러 고루틴이 접근하기 때문에 발생한다.
- 만약 각 고루틴이 같은 자원에 접근하지 않으면 애당초 문제가 발생하지 않는다.

📌 각 고루틴이 서로 다른 자원에 접근하게 만드는 두 가지 방법
- 영역을 나누는 방법
-- 종이의 각 영역을 각자에게 나눠서 서로의 영역을 침범하지 않고 그리면 된다.
-- 각 고루틴은 할당된 작업만 하므로 고루틴 간 간섭이 발생하지 않는다.
-- 그래서 뮤텍스가 필요 없다. 예를 들면 100개 파일을 읽어서 분석할 때는 파일별로 고루틴을 할당해서 수행하면 된다.

- 역할을 나누는 방법
-- 밑그림, 배경 스케치, 채색을 각자 맡아 수행한다면 작업자간 간섭 없이 그림을 완성시킬 수 있다.
-- 이 방법은 채널 개념을 사용해야한다.
*/
