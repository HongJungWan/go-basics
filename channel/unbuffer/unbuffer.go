package main

import "fmt"

func main() {
	ch := make(chan int)       // 1
	ch <- 9                    // 2
	fmt.Println("Never Print") // 3
}

/*
📌 채널 크기
- 일반적으로 채널을 생성하면 크기가 0인 채널이 만들어진다.
- 크기가 0이라는 뜻은 채널에 들어온 데이터를 담아둘 곳이 없다는 말이다.
- 즉, 데이터를 넣을 때 보관할 곳이 없기 때문에 데이터를 빼갈 때까지 대기하게 된다.

예제 코드 시나리오
- 1. 일반적으로 생성하면 크기 0짜리 채널이 만들어진다.
- 2-1. 채널에 데이터를 넣었지만 보관할 곳이 없기 때문에 다른 고루틴에서 데이터를 빼가기를 기다린다.
- 2-2. 하지만 어떤 고루틴도 데이터를 빼가지 않기 때문에
- 3. 실행되지 않고 모든 고루틴이 영원히 대기하게 된다. 따라서 deadlock 메시지를 출력하고 프로그램이 강제 종료된다.
*/
