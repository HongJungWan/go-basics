package main

import "fmt"

// 📌 string 내부 구현
type StringHeader struct {
	Data uintptr
	Len  int
}

func main() {
	str1 := "안녕하세요."

	// str1의 Data와 Len값만 str2에 복사한다.
	// 문자열 주솟값인 Data까지 복사되므로 당연히 문자열 자체는 복사되지 않는다.
	// string 변수끼리 대입 연산에서는 바이트 값만 복사될 뿐 문자열 데이터는 복사되지 않는다.
	str2 := str1

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
📌 UTF-8은 자주 사용되는 영문자, 숫자, 일부 특수 문자를 1바이트로 표현하고 그외 다른 문자들은 2~3바이트로 표현한다.

📌 rune 타입은 Go에서 int32의 별칭으로, 주로 유니코드 문자를 표현할 때 사용된다.

📌 string 타입과 []rune 타입은 상호 타입 변환이 가능하다.
📌 string 타입과 []byte 타입은 상호 타입 변환이 가능하다.

📌 문자열 대소 비교 시, 문자열 길이와 상관없이 앞글자부터 비교한다.
*/
