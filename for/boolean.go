package main

import "fmt"

func main() {
	var a int
	var b int
	found := false

	for a = 1; a < 10; a++ {
		for b = 1; b < 10; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

/*
중첩 for문에서 break를 사용하면 break가 속한 for문에서만 빠져나온다.
모든 for문을 빠져나가고 싶을 때는 첫 번재 방법으로 불리언 변수를 사용하는 것이다.
*/
