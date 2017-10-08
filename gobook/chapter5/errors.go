package main

import (
	"fmt"
	"log"
)

func main() {
	HandleOneError()
}

// HandleOneError typical way to return and handle one type of errors
func HandleOneError() {
	errF := func() (res int, ok bool) {
		res = 2
		ok = true

		return
	}

	res, ok := errF()

	if !ok {
		log.Fatal("Something went wrong")
	}

	fmt.Println(res)
}
