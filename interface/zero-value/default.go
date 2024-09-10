package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker
	att.Attack()
}

/*
📌 인터페이스 변수의 기본값은 유효하지 않은 메모리 주소를 나타내는 nil이다.
- 메모리 주솟값이 nil이면 런타임 오류가 발생한다. 그래서 인터페이스를 사용할 때 항상 인터페이스 값이 nil이 아닌지 확인해야한다.

📌 인터페이스뿐만 아니라 nil 값을 기본으로 갖는 다른 타입 변수 역시 사용하기 전에 nil인지 확인해야 한다.
- 기본값을 nil로 갖는 타입은 포인터, 인터페이스, 함수 타입, 슬라이스, 맵, 채널 등이 있다.
*/
