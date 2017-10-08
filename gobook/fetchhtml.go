package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var url = flag.String("url", "https://golang.org", "url to fetch html response")

func main() {
	flag.Parse()
	flagUrl := *url

	resp, err := http.Get(flagUrl)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Fprint(os.Stdout, string(body))
}
