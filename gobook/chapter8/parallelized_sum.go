package main

import "fmt"

func sum(s []int, c chan<- int) {
	sum := 0

	for i := 0; i < len(s); i++ {
		sum += s[i]
	}

	c <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int)

	go sum(nums[:len(nums)/2], ch)
	go sum(nums[len(nums)/2:], ch)

	r1, r2 := <-ch, <-ch
	fmt.Println(r1, r2, r1+r2)
}
