package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2
	slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)

	fmt.Println(slice)
}

/*
📌 append 함수 시그니처
- func append(slice []Type, elems ...Type) []Type

-- slice []Type: 첫 번째 매개변수로 Type 타입의 슬라이스를 받는다.
-- elems ...Type: 가변 인자(variadic argument)로 Type 타입의 여러 요소를 받을 수 있다.
*/
