package main

import "fmt"

func changeSlice(slice []int) {
	slice[2] = 200
}

func changeArray(array [5]int) {
	array[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)

	fmt.Println("---")

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}

/*
📌 슬라이스와 배열의 동작 차이

- Go 언어에서는 모든 값의 대입은 복사로 일어난다.
- 함수에 Argument(인수)로 전달될 때나 다른 변수에 대입할 때나 값의 이동은 복사로 일어난다.

- 복사는 타입의 값이 복사된다.
-- 포인터는 포인터의 값인 메모리 주소가 복사되고, 구조체가 복사될 때는 구조체의 모든 필드가 복사되고, 배열은 배열의 모든 값이 복사된다.

- Argument(인수)로 전달된 배열은 메모리 공간이 다른, 완전히 다른 배열이기 때문에 changeArray 함수 내에서 값을 변경해도 array 배열은 변경되지 않는다.

- Argument(인수)로 전달된 슬라이스는 sliceHeader 구조체의 각 필드값이 복사되기 때문에 포인터의 메모리 주솟값도 복사되고 len, cap 값도 복사된다.
-- 똑같은 메모리 주솟값을 가지기 때문에 복사된 slice는 같은 배열 데이터를 가리키게 된다.
-- 복사된 slice의 요소를 변경하는 것은 원본 데이터를 직접 변경하는 것과 동일하게 작용한다.

- 결론: 슬라이스 내부에는 배열을 가리키는 포인터가 있고, append()는 슬라이스가 가리키는 배열에 빈 공간이 충분하다면 추가 그렇지 않으면 더 큰 배열을 만들어서 추가한다.
*/
