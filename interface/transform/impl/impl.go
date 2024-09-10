package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Student{15}
	PrintAge(s)
}

/*
📌 인터페이스 변수를 타입 변환을 통해서 구체화된 다른 타입이나 다른 인터페이스로 변환할 수 있다.

- 구체화된 다른 타입으로 타입 변환하기
-- impl.go

- 인터페이스 변수를 다른 구체화된 타입으로 타입 변환할 수 있다.
- 이 방법은 인터페이스를 본래의 구체화된 타입으로 복원할 때 주로 사용한다.
- 사용 방법은 인터페이스 변수에 .(타입)을 붙여 변환한다.
- 핵심은 인터페이스 변수를 본래의 구체적인 타입으로 변환하여 그 타입의 필드에 접근한다.

- 구조체 포인터 *Student 타입은 String() 메서드를 가지고 있다.
- 그래서 PrintAge() 함수의 stringer 인터페이스 인수로 사용할 수 있다.

- Stringer 인터페이스는 String() 메서드만 포함하고 있기 때문에 Stringer 인터페이스 변수로는 Age 값에 접근할 수가 없다.
- stringer 인스턴스 변수 내부에 *Student 타입 인스턴스를 가리키고 있어 *Student 타입으로 에러 없이 변환된다.

- 인터페이스 변수를 구체화된 타입으로 타입 변환하려면 해당 타입이 인터페이스 메서드 집합을 포함하고 있어야 한다.
- 그렇지 않을 경우 컴파일 타임 에러가 발생한다.
*/
