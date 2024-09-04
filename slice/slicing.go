package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println(array)
	fmt.Println(slice)

	array[1] = 100

	fmt.Println("\nAfter Change Element")
	fmt.Println(array)
	fmt.Println(slice)

	slice = append(slice, 500)

	fmt.Println("\nAfter Append 500 Element ")
	fmt.Println(array)
	fmt.Println(slice)
}

/*
📌 슬라이싱(slicing)은 배열의 일부를 집어내는 기능을 말한다.

- array[startIdx:endIdx], 배열의 시작 인덱스부터 끝 인덱스 -1 까지의 배열 일부를 나타내는 슬라이스가 반환된다.
- 주의할 점은 슬라이싱하면 그 결과로 배열 일부를 가리키는 슬라이스를 반환한다.

- 즉, 새로운 배열이 만들어지는게 아니라 배열의 일부를 포인터로 가리키는 슬라이스를 만들어낼 뿐이다.
- ex: 슬라이싱(slicing)을 통해 얻은 슬라이스의 값을 변경하면 원본 배열에도 영향을 미칩니다

📌 슬라이싱 기능은 배열뿐 아니라 슬라이스 일부를 집어낼 때도 사용할 수 있다.
*/
