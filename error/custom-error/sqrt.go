package main

import (
	"fmt"
	"math"
)

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt: %v\n", sqrt)
}

/*
📌 사용자 에러 반환
- fmt.Errorf() 함수를 사용해 에러를 반환했다.
- fmt 패키지의 Errorf() 함수를 이용하면 원하는 에러 메시지를 만들 수 있다.

- errors 패키지의 New() 함수를 이용해서 error를 생성할 수도 있다.
- 인수로 문자열을 입력하면 인수와 같은 메시지를 갖는 error를 생성해서 반환한다.

📌 New() 함수 형식
- func New(text string) error

```
import "errors"

errors.New("에러 메시지")
```
*/
