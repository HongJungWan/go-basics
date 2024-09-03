package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var user = User{name, age}

	// 탈출 검사(Escape Analysis)를 통해 user 변수의 인스턴스가 함수 외부로 공개되는 것을 분석해내서 힙 메모리에 할당
	return &user
}

func main() {
	userPointer := NewUser("홍정완", 27)

	fmt.Println(userPointer)
}

/*
Go 언어는 메모리를 할당할 때 스택 메모리 영역 또는 힙 메모리 영역을 사용한다.
이론상 스택 메모리 영역이 힙 메모리 영역보다 효율적이기 때문에 스택 메모리 영역에서 메모리를 할당하는게 더 좋다.
스택 메모리는 함수 내부에서만 사용 가능한 영역이다. 그래서 함수 외부로 공개되는 메모리 공간은 힙 메모리 영역에서 할당한다.
Go 언어에서는 탈출 검사(Escape Analysis)를 해서 어느 메모리에 할당할지를 결정한다.

함수 외부로 공개되는 인스턴스의 경우 함수가 종료되어도 사라지지 않는다.

- 스택에서 할당된 메모리는 함수의 생명주기와 함께 자동으로 해제되므로 가비지 컬렉터가 관리할 필요가 없다.
*/
