package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b // 함수 리터럴을 사용해서 더하기 함수를 정의하고 반환
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b // 함수 리터럴을 사용해서 곱하기 함수를 정의하고 반환
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4) // 함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result)
}

/*
📌 함수 리터럴(function literal)은 이름 없는 함수로 함수명을 적지 않고 함수 타입 변숫값으로 대입되는 함숫값을 의미한다.
- 함수명이 없기 때문에 함수명으로 직접 함수를 호출할 수 없고 함수 타입 변수로만 호출된다.
- 다른 프로그래밍 언어에서는 익명 함수 또는 람다(Lambda)라고 불린다.
- 다만 Go 언어에서는 `함수 리터럴`이라고 부른다.
*/
