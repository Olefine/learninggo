package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	time := time.Tick(1 * time.Second)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for x := 10; x != 0; x-- {
		select {
		case <-time:
			fmt.Println(x)
		case <-abort:
			fmt.Println("aborting")
			return
		}
	}

	fmt.Println("GO!")
}
