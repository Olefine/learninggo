package main

import "os"
import "io/ioutil"
import "fmt"
import "strings"

func main() {
	fileNames := os.Args[1:]
	counts := make(map[string]int)

	for _, fileName := range fileNames {
		value, err := ioutil.ReadFile(fileName)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Something went wrond %v\n", err)
			continue
		}

		// need to use string() because value is []byte
		for _, line := range strings.Split(string(value), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		fmt.Println(line, ":", n)
	}
}