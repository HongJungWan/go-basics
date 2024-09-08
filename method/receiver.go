package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	// 포인터 메서드 테스트
	var a *account = &account{100, "Joe", "Park"}
	a.withdrawPointer(30)
	fmt.Println(a.balance)

	a.withdrawValue(20)
	fmt.Println(a.balance)

	// 값 타입 메서드 테스트
	var b account = account{100, "Joe", "park"}
	b.withdrawValue(30)
	fmt.Println(b.balance)

	result := b.withdrawReturnValue(30)
	fmt.Println(result.balance)
}

/*
📌 포인터 메서드 vs 값 타입 메서드
-- 포인터 메서드를 호출하면 포인터가 가리키고 있는 메모리의 주솟값이 복사된다.
-- 반면 값 타입 메서드를 호출하면 리시버 타입의 모든 값이 복사된다. 예를 들어 리시버 타입이 구조체이면 구조체의 모든 데이터가 복사된다.


📌 언제 포인터 메서드를 만들고 값 타입 메서드를 만들어야 할까?
-- 포인터 메서드는 메서드 내부에서 리시버의 값을 변경시킬 수 있다.
-- 값 타입 메서드는 호출하는 쪽과 메서드 내부의 값은 별도 인스턴스로 독립되기 때문에 메서드 내부에서 리시버의 값을 변경시킬 수 없다.
-- 정리하면 포인터 메서드는 인스턴스 중심이고 값 타입 메서드는 값 중심이 된다.
*/
