package main

import (
	"mymath"
	"fmt"
)

const (
	y = 1
	z = 2
	h = 3
	x = iota
)

func main() {
	var ar = [10]int{1, 2, 3, 4, 5}
	var a, b []int
	a = ar[1:3]
	b = ar[3:5]
	for index := 0; index < len(a); index++ {
		fmt.Println(index)
		fmt.Println(a[index])
		fmt.Println(b[index])
	}

	fmt.Printf("hello , world %d", x)
	fmt.Printf("hello, world . Sqrt(2) = %v \n", mymath.Sqrt(2))
}
