package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("err", err)
		return
	}

	defer fmt.Println("반드시 호출된다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Hello World를 쓴다.")
	fmt.Fprintln(f, "Hello World") // 파일에 텍스트를 쓴다.
}

/*
- 때론 함수가 종료되기 직전에 실행해야 하는 코드가 있을 수 있다.
- 대표적으로 파일이나 소켓 핸들처럼 OS 내부 자원을 사용하는 경우다.
- 함수 종료 전에 처리해야 하는 코드가 있을 때 defer를 사용해 실행할 수 있다.

-- `defer 명령문`
-- 위와 같이 적으면 명령문이 바로 실행되는 게 아닌 해당 함수가 종료되기 직전에 실행되도록 지연된다.
-- defer는 역순으로 호출된다. (아래부터 호출)
*/
