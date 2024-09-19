package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
			fmt.Println("고루틴 끝")
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative vlaue: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

/*
📌 동시성 프로그래밍 주의점
- 동시성 프로그래밍의 문제점은 동일한 메모리 자원에 여러 고루틴이 접근할 때 발생한다.
- 고루틴은 각 CPU 코어에서 별도로 동작하지만 같은 메모리 공간에 동시에 접근해서 값을 변경시킬 수 있다.

📌 문제 원인
- account.Balance += 1000 코드에 있다. 이 코드는 먼저 Balance 값을 읽고 1000을 더해서 Balance에 다시 저장하는 두 단계로 이뤄진다.
- 이때 첫 번째 단계가 완료 되기 전에 다른 고루틴이 첫 번째 단계를 수행하면 두 고루틴은 똑같은 값을 읽어서 1000씩 더해 다시 Balance에 저장한다.
- 즉, 고루틴 2개가 각각 입금을 했는데 한 번 입금한 효과밖에 나지 않게 된다. 이 상태에서 출금이 각각 이루어지면 잔고가 -가 된다.

📌 뮤텍스를 이용한 동시성 문제 해결
- 가장 단순한 해결 방법은 한 고루틴에서 값을 변경할 때 다른 고루틴이 건들지 못하게 하는 것이다.
- 뮤텍스를 이용하면 자원 접근 권한을 통제할 수 있다.
- 뮤텍스는 mutual exclusion의 약자이다.

- 뮤텍스의 Lock() 메서드를 호출해 뮤텍스를 획득할 수 있다.
- 이미 Lock() 메서드를 호출해서 다른 고루틴이 뮤텍스를 획득했다면 나중에 호출한 고루틴은 앞서 획득한 뮤텍스가 반납될 때까지 대기하게 된다.
- 사용 중이던 뮤텍스는 Unlock() 메서드를 호출해서 반납한다.
- 이후 대기하던 고루틴 중 하나가 뮤텍스를 획득한다.

- 한 번 획득한 뮤텍스는 반드시 Unlock()을 호출해서 반납해야 한다.
*/
