package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	n := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(n)
	fmt.Println(n, "\n")

	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
		{"켄", 38}, {"송하나", 18},
	}

	sort.Sort(Students(s))
	fmt.Println(s)
}

/*
📌 func Ints(x []int) { Sort(IntSlice(x)) }
- Ints sorts a slice of ints in increasing order.

📌 func Sort(data Interface)
- Sort sorts data in ascending order as determined by the Less method.

📌 Students(s)
- []Student 타입인 s를 정렬 인터페이스를 포함한 타입인 Students 타입으로 변환하는 구문이다.
- []Student 타입은 정렬에 필요한 Len(), Less(), Swap() 메서드를 가지고 있지 않기 때문에,
   sort.Sort() 함수의 인수로 사용될 수 없다. 그래서 []Student의 별칭 타입을 만들어서 정렬 인터페이스를 포함한다.
*/
