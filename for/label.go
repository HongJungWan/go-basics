package main

import "fmt"

func main() {
	var a int
	var b int

OuterFor:
	for a = 1; a < 10; a++ {
		for b = 1; b < 10; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

/*
플래그 변수가 아닌 레이블을 사용해서 for문을 종료할 수도 있다.

for문을 시작할 때 레이블을 정의하고 break할 때 앞서 정의한 레이블을 적어주면,
그 레이블에서 가장 먼저 속한 for문까지 모두 종료한다.
*/
