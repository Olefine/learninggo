package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4}

	fmt.Println(s)
	f(s)
	fmt.Println(s)
}

func f(a []int) {
	a = []int{}

	fmt.Println(a)
}
