package main

import (
	"fmt"
	"os"
)

type Writer func(string) // 함수 타입을 정의하여 string을 인수로 받는 함수를 Writer 타입으로 선언

func writeHello(writer Writer) { // Writer 타입의 함수를 인수로 받는 writeHello 함수
	writer("Hello World") // Writer 타입 함수 호출, 내부 동작은 외부에서 주입된 함수에 따라 결정됨
}

func main() {
	file, err := os.Create("test.txt") // 파일을 생성
	if err != nil {
		fmt.Println("Failed to create a file") // 파일 생성 실패 시 메시지 출력 후 함수 종료
		return
	}

	defer file.Close() // main 함수 종료 시 파일을 닫음

	// 의존성 주입: writeHello에 익명 함수를 전달하여, 메시지를 파일에 쓰는 동작을 구현
	writeHello(func(msg string) {
		fmt.Fprintln(file, msg) // 파일에 메시지를 출력
	})
}

/*
📌 파일 핸들을 내부 상태로 사용하는 예
- 함수 리터럴을 이용해서 원하는 함수를 그때 그때 정의해서 함수 타입 변숫값으로 사용할 수 있다.
- 또 필요한 외부 변수를 내부 상태로 가져와서 편리하게 사용할 수 있다.

📌 의존성 주입
- writeHello() 함수 입장에서 생각해보자.
- writeHello()는 인수로 Writer 함수 타입을 받는다.
- writeHello() 함수 입장에서 보면 인수로 온 writer를 호출했을 때 그게 파일에 씌여질지 네트워크로 전송될지, 프린터로 찍힐지
- 아무 상관없는 로직이 수행될지 알 수 없다.

- 이렇게 외부에서 로직을 주입하는 것을 의존성 주입이라고 한다.
*/
