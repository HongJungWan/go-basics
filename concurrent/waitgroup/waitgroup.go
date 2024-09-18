package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)

	wg.Done()
}

func main() {
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go SumAtoB(1, 100000000000)
	}

	wg.Wait()
}

/*
📌 서브 고루틴이 종료될 때까지 기다리기
- main() 고루틴이 몇초 대기하는 것으로 모든 실행을 보장했다. 그러나 늘 함수 실행 시간을 알 수 있는 것은 아니다. 그렇다면 어떻게 고루틴이 종료될 때까지 대기할 수 있을가?
- 바로 sync 패키지의 WaitGroup 객체를 사용하면 된다.

```
var wg sync.WaitGroup

wg.Add(3)
wg.Done()
wg.Wait()
```

Add() 메서드를 통해 완료해야 하는 작업 개수를 설정하고 각 작업이 완료될 때마다 Done() 메서드를 호출하여 남은 작업 개수를 하나씩 줄여준다.
Wait()는 전체 작업이 모두 완료될 때까지 대기한다.
*/
