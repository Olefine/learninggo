package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
}

func generateError() error {
	err := errors.New("Hello from error")

	return err
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		fmt.Println("retry")
		_, err := http.Head(url)
		if err == nil {
			return nil
		}

		log.Printf("server not respoding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)

}
