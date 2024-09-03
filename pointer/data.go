package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	//var p1 *Data = &Data{}
	//var p2 *Data = p1
	//var p3 *Data = p1

	//p1 := &Data{}      // &를 사용하는 초기화
	//var p2 = new(Data) // new()를 사용하는 초기화
}

/*
📌 변수 대입이나 함수 인수 전달은 항상 값을 복사하기 때문에 많은 메모리 공간을 사용하는 문제가 존재한다.
📌 또한 다른 공간으로 복사되기 때문에 변경 사항이 적용되지 않는다.

📌 인스턴스는 메모리에 존재하는 데이터의 실체다. -> 포인터를 이용해서 인스턴스에 접근할 수 있다.
📌 쓸모 없어진 인스턴스는 가비지 컬렉터가 자동으로 지워준다.
*/
