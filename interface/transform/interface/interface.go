package main

func main() {

}

/*
📌 인터페이스 변수를 타입 변환을 통해서 구체화된 다른 타입이나 다른 인터페이스로 변환할 수 있다.

- 다른 인터페이스로 타입 변환하기
-- interface.go

- 인터페이스 변환을 통해 구체화된 타입뿐 아니라 다른 인터페이스로 타입 변환할 수 있다.
- 이때는 구체화된 타입으로 변환할 때와는 달리 변경되는 인터페이스가 변경 전 인터페이스를 포함하지 않아도 된다.
- 하지만 인터페이스가 가리키고 있는 실제 인스턴스가 변환하고자 하는 다른 인터페이스를 포함해야한다.

- 예를 들어,

var a AInterface = ConcreteType{}
b := a.(BInterface)

- ConcreteType이 AInterface와 BInterface 인터페이스를 모두 포함하고 있는 경우에
- 위와 같이 ConcreteType 인스턴스를 가리키고 있는 AInterface 변수 a는 BInterface로 타입 변환이 가능하다.
- 그 이유는 a가 가리키고 있는 ConcreteType 인스턴스는 BInterface도 포함하고 있기 때문이다.


📌 타입 변환 성공 여부 반환
var a Interface
t, ok := a.(ConcreteType)


📌 타입 변환 결과를 반환받아서 변환 성공 여부를 검사하는 if문을 살펴봅시다.

c, ok := reader.(Closer)
if ok {
	...
}

if c, ok := reader.(Closer); ok {
	...
}
*/
