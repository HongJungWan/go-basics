package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)

		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)

		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)

	// 포크와 수저 뮤텍스
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")

	wg.Wait()
}

/*
📌 뮤텍스와 데드락
- 뮤텍스를 사용하면 동시성 프로그래밍 문제를 해결할 수 있지만, 또 다른 문제가 발생할 수 있다.
- 첫 번째 문제는 동시성 프로그래밍으로 얻는 성능 향상을 얻을 수 없다는 점이다.
- 뮤텍스는 오직 하나의 고루틴만 공유 자원에 접근할 수 있도록 제한한다. 따라서 여러 고루틴 중 뮤텍스를 획득한 고루틴만 실행된다.
- 성능을 향상시키려고 동시성 프로그램을 구현했지만 성능 향상을 얻지 못하는 아이러니한 문제가 생기게 된다.

- 두 번재 문제는 데드락이 발생할 수 있다는 점이다.
- 데드락은 프로그램을 완전히 멈추게 만들어 버린다.
- 어떤 고루틴도 원하는 만큼 뮤텍스를 확보하지 못한다면 무한히 대기하게 되는 경우를 데드락이라고 말한다.
*/

/*
- 위 예를 보면 A는 포크 뮤텍스를 획득했고, B는 수저 뮤텍스를 획득했지만 서로 두 번째 뮤텍스를 획득하지 못해 무한히 대기하게 되고, Go 언어에서는 데드락을 감지하고 에러를 반환한다.
- 둘 다 수저 먼저 잡고 포크를 집게 만들면 프로그램이 멈추는 문제는 발생하지 않는다.
- 하지만 실제 실무에서는 뮤텍스들이 복잡하게 꼬여 있어서 단순히 순서만 변경해서는 해결할 수 없는 경우가 있다.
- 그래서 데드락 문제는 동시성 프로그래밍에서 해결하기 힘든 난제다.
*/

/*
📌 데드락이 발생하는 이유 - 정리
- 두 개의 고루틴(A와 B)이 두 개의 락(fork와 spoon)을 서로 다른 순서로 획득하려고 하면 데드락(deadlock) 상황이 발생할 수 있다.

- 데드락은 동시성 프로그래밍에서 두 개 이상의 스레드(Go에서는 고루틴)가 서로 다른 스레드가 소유하고 있는 자원을 기다리느라 영원히 멈추는 상태를 말한다.
- Go에서는 두 개 이상의 고루틴이 동일한 뮤텍스를 다른 순서로 획득하려고 할 때 주로 데드락이 발생한다.
*/
