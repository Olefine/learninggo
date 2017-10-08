package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	boom := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println("   .")
			time.Sleep(1 * time.Second)
		}
	}
}
