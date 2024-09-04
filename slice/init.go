package main

import "fmt"

func main() {
	// 📌 배열 선언
	var array1 = [3]int{1, 2, 3}
	var array2 = [10]int{1, 0, 0, 0, 0, 2, 0, 0, 0, 3}
	fmt.Println("\n", array1, "\n", array2)

	// 📌 슬라이스 선언
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}
	var slice3 = make([]int, 3)

	slice3[1] = 5

	for i, item := range slice3 {
		slice3[i] = item + 1
	}

	fmt.Println("\n", slice1, "\n", slice2, "\n", slice3)
}

/*
📌 슬라이스를 초기화하지 않으면 길이가 0인 슬라이스가 만들어진다.
*/
