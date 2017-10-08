package main

import "fmt"
import "os"
import "bufio"

func main() {
	findDuplicatedLines()
}

func commandLineArgs() {
	for _, value := range os.Args[1:] {
		fmt.Println(value)
	}
}

func findDuplicatedLines() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Println("Line -", line, "n -", n)
	}
}
