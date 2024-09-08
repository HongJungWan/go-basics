package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		if i == 6 {
			break
		}

		fmt.Println("6 * ", i, "=", 6*i)
	}
}

/*
# 📌 `기본`
for 초기문; 조건문; 후처리 {
	코드 블록 (조건문이 true인 경우 수행된다.)
}


# 📌 `초기문 생략`
for ; 조건문; 후처리 {
	코드 블록
}


# 📌 `후처리 생략`
for 초기문; 조건문; {
	코드 블록
}


# 📌 `조건문만 사용`
for ; 조건문; {
	코드 블록
}

for 조건문; {
	코드 블록
}


# 📌 `무한 루프`
for true {
	코드 블록
}


# 📌 `정숫값으로 순회`
for i := range 10 {
	fmt.Println(i)
}


📌 continue와 break
- continue와 break는 반복문을 제어하는 키워드다.
*/
