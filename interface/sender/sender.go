package main

import (
	"go-basics/interface/fedex"
	"go-basics/interface/post"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	sender1 := &fedex.FedexSender{}
	SendBook("GSP", sender1)
	SendBook("HABIB", sender1)

	println()

	sender2 := &post.PostSender{}
	SendBook("GSP", sender2)
	SendBook("HABIB", sender2)
}

/*
- 함수 인수로 인터페이스를 전달.

- 내부 동작을 감춰서 서비스를 제공하는 쪽과 사용하는 쪽 모두에게 자유를 주는 방식을 추상화라고 한다.
- 인터페이스는 추상화를 제공하는 추상화 계층이다.

- 위 예제에서 Sender 인터페이스라는 추상화 계층을 만듦으로써 택배 서비스를 사용하는 쪽도 택배 서비스를 제공하는 쪽도
- 서로 구현에 신경 쓰지 않고 추상화된 관계 중심으로 코딩할 수 있다.

- 이렇게 추상화 계층을 이용해 의존 관계를 끊는 것을 디커플링(decoupling)이라 말한다.
*/
