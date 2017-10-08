package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			numbers <- x
		}
	}()

	go func() {
		for x := range numbers {
			squares <- x * x
		}
		close(squares)
	}()

	for x := 0; x < 10; x++ {
		fmt.Println(<-squares)
	}
}
