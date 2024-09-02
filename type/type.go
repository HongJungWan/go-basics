package main

import "fmt"

func main() {
	var b float64 = 3.5
	var g int = int(b * 3)
	var h int = int(b)
	fmt.Printf("g: %d, h: %d", g, h)

	fmt.Println("\n")

	var a int16 = 3456
	var c int8 = int8(a)
	fmt.Printf("a: %d, c: %d", a, c)
}

/*
📌 실수 타입에서 정수 타입으로 타입 변환하면 소수점 이하 숫자가 없어진다.
📌 큰 범위를 갖는 타입에서 작은 범위를 갖는 타입으로 변화하면 값이 달라질 수 있다.
*/
