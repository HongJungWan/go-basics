package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int // int 타입 인수 2개를 받아서 int 타입을 반환하는 함수 타입 변수
	operator = getOperator("*")

	var result = operator(3, 4) // 함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result)
}

/*
📌 함수 타입 변수
- 함수 타입 변수란 함수를 값으로 갖는 변수를 의미한다.
- 어떻게 함수를 값으로 가질 수 있을까? -> 컴퓨터는 0,1로 나타낼 수 있는 숫자값만 가질 수 있고 포인터는 숫자로 나타낼 수 있는 메모리 주소를 값으로 가진다.
- 즉, 함수를 숫자로 표현하면 된다.

📌 기본기
- 초창기 프로그램은 천공카드의 연속, 즉 명령어를 적은 긴 종이 다발로 볼 수 있다.
- 이 종이 다발에는 프로그램 코드가 잔뜩 적혀 있고 함수 역시 코드 블록이기 때문에 이 종이 다발에 적혀 있다.

- 1번 라인에서 main() 함수가 시작되고 100번 라인에서 f() 함수가 시작된다고 가정하자.
- 컴퓨터 내부(정확히는 CPU 내부)에는 프로그램 카운터(Program Counter)가 있다.

--- 프로그램 카운터는 다음 실행할 라인을 나타내는 레지스터다.
--- 1번 라인 명령을 실행하면 프로그램 카운터는 1이 증가하여 2를 가리키고 다음에 2번 라인을 실행하게 된다.
--- 만약 f() 함수가 호출되면 프로그램 카운터는 f() 함수의 시작 포인트 즉, 100번 라인으로 변경되어서 다음 번에 100번 라인부터 명령을 실행하게 된다.
--- 이것이 바로 함수 호출 과정이다.

- 즉, 함수 시작 지점 역시 숫자로 표현할 수 있다.
- 이 함수 시작 지점이 바로 함수를 가리키는 값이고, 마치 포인터처럼 함수를 가리킨다고 해서 함수 포인터(function pointer)라고 부른다.
*/

/*
📌 별칭으로 함수 정의 줄여 쓰기
- 함수 정의는 일반적으로 길다. 따라서 매번 함수 정의를 쓰면 코드 가독성이 떨어진다.
- 이럴 때는 별칭 타입을 써서 함수 정의를 짧게 줄일 수 있다.

```
type opFunc func (int, int) int
```
- 위와 같이 func (int, int) int 함수 정의를 opFunc으로 짧게 재정의하면 getOperator() 함수 정의도 짧게 적을 수 있다.


```
func getOperator(op string) opFunc
```
- 위와 같이 func (int, int) int 함수 정의를 opFunc으로 짧게 재정의하면 getOperator() 함수 정의도 짧게 적을 수 있다.
*/

/*
📌 함수 정의에서 매개변수명 생략하기
- 함수 정의에서 매개변수명은 적어도 되고 적지 않아도 된다.

```
func (int, int) int
```

라고 해도 되고

```
func (a int, b int) int
```

라고 해도 된다.
*/
