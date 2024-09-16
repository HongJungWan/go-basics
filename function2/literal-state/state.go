package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3) // 함수 리터럴 3개를 가진 슬라이스
	fmt.Println("Value Loop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]() // f[i]에 저장된 함수를 호출하는 의미
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("Value Loop2")
	for i := 0; i < 3; i++ {
		v := i // 🚩 v 변수에 i값 복사, i 변수가 참조로 캡쳐되는 문제 현상을 값 복사로 해결

		f[i] = func() {
			fmt.Println(v) // 캡쳐된 v값 출력
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}

/*
📌 함수 리터럴 내부 상태.
- 함수 리터럴은 필요한 변수를 내부 상태로 가질 수 있다.
- 함수 리터럴 내부에서 사용되는 외부 변수는 자동으로 함수 내부 상태로 저장된다.

func main() {
	i := 0

	f := func() {
		i += 10 // i에 10 더하기
	}

	i++

	f() // f 함수 타입 변수를 사용해서 함수 리터럴 실행

	fmt.Println(i) // 11
}

📌 i값이 함수 리터럴이 정의되는 시점이 아닌 함수가 호출되는 시점 값으로 사용되는 것을 주의하자.
- 함수 리터럴이 정의되는 시점에 i는 0이지만 함수가 호출되는 시점은 i++이 실행된 다음이기 때문에 리터럴 내부의 i는 0이 아니라 1이 된다.
- 함수 리터럴에서 외부 변수를 내부 상태로 가져올 때 값 복사가 아닌 인스턴스 참조로 가져오게 된다.
- 포인터 형태로 가져온다고 보면 편하다.


📌 함수 리터럴 외부 변수를 내부 상태로 가져오는 캡쳐(capture)라고 한다.
- 캡쳐는 값 복사가 아닌 참조 형태로 가져오게 되니 주의해야 한다.


📌 참조로 변수를 가져온다는게 무엇일까?
- 변수의 주소를 포인터 값으로 복사한다고 보면 된다.
- 함수 리터럴 외부의 변수 i를 캡쳐할 때 변수 i의 주솟값을 포인터 형태로
  함수 리터럴 내부 상태로 가져와 나중에 캡쳐된 내부 상태를 사용할 때 메모리 주솟값을 통해 외부 변수 i에 접근하게 된다.

- 예를 들어 고루틴에 의해서 함수 리터럴이 여러 고루틴에서 실행될 때 주의해야 한다.
*/
