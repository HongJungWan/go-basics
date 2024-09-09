package main

import "fmt"

type StudentA struct {
	Age int
}

func PrintVal(v interface{}) {
	// Type Assertion은 인터페이스에 저장된 구체적인 값의 타입을 가져온다.
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not Supported Type: %T:%v\n", t, t)
	}
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(StudentA{15})
}

/*
- interface{}는 메서드를 가지고 있지 않은 빈 인터페이스다.
- 가지고 있어야 할 메서드가 하나도 없기 때문에 모든 타입이 빈 인터페이스로 쓰일 수 있다.
- 빈 인터페이스는 어떤 값이든 받을 수 있는 함수, 메서드, 변숫값을 만들 때 사용한다.
*/
