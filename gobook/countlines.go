package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Println("Line -", line, "n -", n)
	}
}

func countLines(in *os.File, store map[string]int) {
	input := bufio.NewScanner(in)

	for input.Scan() {
		store[input.Text()]++

		if store[input.Text()] > 1 {
			fmt.Println(in.Name())
		}
	}
}