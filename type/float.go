package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123

	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a) // 1234.523
	fmt.Println(b) // 3456.123
	fmt.Println(c) // 4.266663e+06 --> 4266663.334329
	fmt.Println(d) // 1.2799989e+07 --> 12799990.002987
}

/*
📌 실수는 무한히 많은 숫자가 있기 때문에 최솟값, 최댓값보다는 표현할 수 있는 소수부 자릿수에 더 주의를 기울여야 한다.
📌 일반적으로 float32는 7자리까지 표현, float64는 15자리까지 표현

📌 Go에서는 실숫값을 표현할 때 정확한 값이 아닌 타입이 허용하는 범위에서 가장 가까운 근삿값으로 표현됨.
📌 위 예제 코드처럼 실수 타입 잘못 사용하면 오차가 발생함.
*/
