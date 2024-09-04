package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}

	slice2 := append(slice, 4)
	fmt.Println(slice2)

	slice2 = append(slice, 1, 2, 3, 4)
	fmt.Println(slice2)
}
