package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

const dirPath string = "/home/egorodov"

func main() {
	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		log.Fatalf("Cannot open dir: %v", err)
	}

	var filtered []string

	for _, fileInfo := range files {
		filename := fileInfo.Name()
		if strings.HasSuffix(filename, ".csv") {
			filtered = append(filtered, dirPath+"/"+filename)
		}
	}

	filenamesChan := make(chan string)
	go func(ch chan string) {
		for _, f := range filtered {
			ch <- f
		}

		defer close(ch)
	}(filenamesChan)

	fmt.Println("The total size is", collectFiles(filenamesChan))
}

func collectFiles(filenames <-chan string) int64 {
	sizes := make(chan int64)

	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)

		go func(fl string) {
			defer wg.Done()

			file, err := os.Open(fl)
			if err != nil {
				log.Fatalf("%v", err)
			}
			defer file.Close()

			filestat, _ := file.Stat()
			filesize := filestat.Size()
			fmt.Println(filesize, "-", fl)

			sizes <- filesize
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}
