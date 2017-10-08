package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const baseHost = "http://numbersapi.com/"

func main() {
	reader := make(chan string)

	for i := 1; i < 90; i++ {
		go readUrl(i, reader)
	}

	for i := 1; i < 90; i++ {
		fmt.Println(<-reader)
	}

	close(reader)
}

func readUrl(number int, out chan<- string) {
	c := http.Client{}
	urlToFetch := baseHost + strconv.Itoa(number)
	resp, err := c.Get(urlToFetch)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	res, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	out <- string(res)
}
