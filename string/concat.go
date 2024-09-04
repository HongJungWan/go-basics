package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var rst string

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder

	for _, c := range str {
		if c >= 'a' && c <= 'Z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	str := "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}

/*
📌 Go 언어는 기존 문자열 메모리 공간을 건드리지 않고,
📌 새로운 메모리 공간을 만들어서 두 문자열을 합치기 때문에 string 합 연산 이후 주솟값이 변경된다. -> 문자열 불변 원칙이 준수된다.

📌 string 합 연산을 빈번하게 하면 메모리가 낭비된다.
📌 그래서 string 합 연산을 빈번하게 사용하는 경우에는 strings 패키지의 Builder를 이용해서 메모리 낭비를 줄일 수 있다.

📌 Go 언어는 빈번한 합 연산 시, 메모리가 낭비되는 데도 왜 문자열 불변 원칙을 지키려 할까?
📌 가장 큰 이유는 예기치 못한 버그를 방지하기 위해서다.
📌 Go 언어에서 문자열이 불변인 이유는 문자열이 변경될 수 있다면, 해당 문자열을 가리키는 모든 참조가 영향을 받을 수 있기 때문이다.
*/
