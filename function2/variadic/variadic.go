package main

import "fmt"

func sum(numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(sum(1))
}

/*
📌 함수 인수 개수가 고정적이지 않은 함수를 가변 인수 함수(variadic function)이라고 한다.

- ... 키워드를 사용해서 가변 인수를 처리할 수 있다.
- 인수 타입 앞에 ...를 붙여서 해당 타입 인수를 여러 개 받는 가변 인수임을 표시하면 된다.
- 가변 인수는 함수 내부에서 해당 타입의 슬라이스로 처리된다.
*/
