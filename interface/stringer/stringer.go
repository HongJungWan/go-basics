package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	age  int
}

func (s Student) String() string {
	return fmt.Sprintf("%d, %s", s.age, s.Name)
}

func main() {
	s := Student{"hong", 15}

	var stringer Stringer
	stringer = s

	fmt.Println(stringer.String())
}

/*
- 인터페이스란 구현을 포함하지 않는 메서드 집합이다.
- 구체화된 타입이 아닌 인터페이스만 가지고 메서드를 호출할 수 있다.
- Go 언어에서는 인터페이스 구현 여부를 그 타입이 인터페이스에 해당하는 메서드를 가지고 있는지로 판단하는 덕 타이핑을 지원한다.
*/
