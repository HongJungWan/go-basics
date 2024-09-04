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
}

/*
📌 슬라이스와 배열의 동작 차이

- Go 언어에서는 모든 값의 대입은 복사로 일어난다.
- 함수에 Argument(인수)로 전달될 때나 다른 변수에 대입할 때나 값의 이동은 복사로 일어난다.


*/
