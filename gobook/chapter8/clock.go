package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

type request struct {
	Name string    `json: "name"`
	Time time.Time `json: "time"`
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		req := &request{Name: "Avails #1", Time: time.Now()}
		err := json.NewEncoder(c).Encode(req)

		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
