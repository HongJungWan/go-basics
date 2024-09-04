package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)
}

/*
📌 copy() : 슬라이스 복사, 반환 값은 실제로 복사된 요소 개수
*/
