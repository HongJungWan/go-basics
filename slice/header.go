package main

import "fmt"

// 📌 slice 내부 구현
type sliceHeader struct {
	Data uintptr // 실제 배열을 가리키는 포인터
	Len  int     // 요소 개수
	Cap  int     // 실제 배열의 길이
}

func main() {
	// len 3, cap 3인 slice
	slice1 := make([]int, 3)
	fmt.Println(slice1)

	// len 3, cap 5인 slice
	// 현재는 3개만 사용하지만 나머지 추가될 요소를 위해 비어두는 경우.
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)
}

/*
📌 슬라이스는 실제 배열을 가리키는 포인터를 가지고 있다.
📌 이로 인해 다른 배열을 가리키도록 변경할 수 있고, 슬라이스는 배열에 비해서 사용되는 메모리나 속도에 이점이 있다.

- 슬라이스는 포인터를 통해 배열의 특정 부분을 참조할 수 있기 때문에, 배열 전체를 복사하지 않고도 동일한 데이터를 여러 슬라이스가 참조할 수 있다.
- 즉, 슬라이스 간의 복사는 배열 전체를 복사하는 것이 아니라, 슬라이스 구조체 자체만 복사한다.

- 반면 배열은 크기가 고정되어 있으며, 배열 간의 대입은 배열의 모든 요소를 복사한다.
*/
