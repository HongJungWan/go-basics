package main

import "fmt"

func divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}

	result = a / b
	success = true
	return
}

func main() {
	c, success := divide(9, 3)
	fmt.Println(c, success)
}

/*
📌 인수(Argument)는 매개변수(Parameter)로 복사된다.
📌 함수 선언부에 반환 타입을 적을 때 변수명까지 지정해주면, return 문으로 해당 변수를 명시적으로 반환하지 않아도 값을 반환할 수 있다.
*/
