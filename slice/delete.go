package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	//for i := idx + 1; i < len(slice); i++ {
	//	slice[i-1] = slice[i]
	//}
	//slice = slice[:len(slice)-1]

	slice = append(slice[:idx], slice[idx+1:]...)

	fmt.Println(slice)
}

/*
📌 요소 삭제 : 중간 요소를 삭제하고, 중간 요소 이후의 값을 앞당겨서 삭제된 요소를 채운다.
*/
