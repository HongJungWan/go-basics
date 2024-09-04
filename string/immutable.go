package main

import "fmt"

func main() {
	str := "Hello World"
	str = "Hello Golang" // 전체 바꾸기는 가능
	// str[2] = 'a'         // Error 발생, 일부 바꾸기는 불가능
	fmt.Println(str)

	str1 := "Hi World"
	slice := []byte(str1) // 타입 변환할 때 문자열을 복사해서, 새로운 메모리 공간을 만들어 슬라이스가 가리키도록 한다.
	slice[2] = 'a'
	fmt.Printf("%s\n", slice)
}

/*
📌 문자열은 불변(immutable)이다.
📌 불변이라는 말은 string 타입이 가리키는 문자열의 일부만 변경할 수 없다는 말이다.
*/
