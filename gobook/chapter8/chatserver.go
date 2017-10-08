package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

type Room []net.Conn

var room = Room{}

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		log.Fatalf("Problem with connection: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	room = append(room, conn)
	conn.Write([]byte("Connected to the room\n"))

	for {
		received, _ := ioutil.ReadAll(conn)
		receivedValue := string(received)
		fmt.Println(receivedValue)
		sendTo(conn, receivedValue)
	}

}

func sendTo(from net.Conn, msg string) {
	for _, dst := range room {
		if dst != from {
			dst.Write([]byte(msg))
		}
	}
}
