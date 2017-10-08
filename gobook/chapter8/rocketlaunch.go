package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Starting the launch...")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("Aborted")
		return
	}

	launch()
}

func launch() {
	fmt.Println("Launchiin")
}
