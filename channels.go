package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	closingChannels()
}

func simpleGotoutines() {
	printV := func(from string) {
		for i := 0; i <= 3; i++  {
			fmt.Println(from, ":", i)
		}
	}

	printV("direct")
	go printV("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("ongoing")

	fmt.Scanln()
	fmt.Println("done")
}

func simpleChannel() {
	messages := make(chan string)

	go func() { messages <- "Ping"}()

	msg := <- messages

	fmt.Println(msg)
}

func channelBuffering() {
	messages := make(chan string)

	messages <- "Hello"
	messages <- "World"

	fmt.Println(<- messages)
	fmt.Println(<- messages)
}

//Block until someone send data to the channel
func channelSync() {
	channel := make(chan bool, 1)

	go func(bchan chan bool) {
		fmt.Println("Starting")
		time.Sleep(time.Second)
		fmt.Println("done")

		bchan <- true
	}(channel)

	<- channel
}

//1 function send data to the channel
//2 function recieve data from first channel and then put this message to channel 2
//main function receives data from channel 2
func channelsDirection() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping := func(out chan <- string, message string) {
		out <- message
	}

	pong := func(in <- chan string, out chan <- string) {
		msg := <- in
		out <- msg
	}

	ping(pings, "ping")
	pong(pings, pongs)

	fmt.Println(<- pongs)
}

func multiplexor() {
	firstSource := make(chan string, 1)
	secondSource := make(chan string, 1)

	outStream := make(chan string)

	go func(){
		for {
			select {
			case msg := <-firstSource:
				outStream <- msg
			case msg := <-secondSource:
				outStream <- msg
			}
		}
	}()

	go func(out chan <- string) {
		for i := 0; i < 10; i++ {
			out <- "firstSource: " + strconv.Itoa(i)
			time.Sleep(time.Second)
		}
	}(firstSource)

	go func(out chan <- string) {
		for i := 0; i < 10; i++ {
			out <- "secondSource :" + strconv.Itoa(i)
			time.Sleep(time.Second)
		}
	}(secondSource)

	for {
		select {
		case msg := <-outStream:
			fmt.Println(msg)
		}
	}
}

func from2Channels() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(out chan <- string) {
		time.Sleep(time.Second)
		out <- "First"

	}(c1)

	go func(out chan <- string) {
		time.Sleep(time.Second)
		out <- "Second"
	}(c2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <- c1:
			fmt.Println(m1)
		case m2 := <- c2:
			fmt.Println(m2)
		}
	}
}

func timeouts() {
	c1 := make(chan string)

	go func(out chan <- string) {
		time.Sleep(time.Second * 2)
		out <- "Channel 1"
	}(c1)

	select {
		case msg := <- c1:
			fmt.Println(msg)
		case <- time.After(time.Second * 1):
			fmt.Println("Timeout for channel 1")
	}

	c2 := make(chan string)

	go func(out chan <- string) {
		time.Sleep(time.Second * 2)
		out <- "Channel 2"
	}(c2)

	select {
		case msg := <- c2:
			fmt.Println(msg)
		case <- time.After(time.Second * 3):
			fmt.Println("Timeout for channel 2")
	}
}

func nonBlockingSend() {
	messages := make(chan string)

	select {
	case msg := <- messages:
		fmt.Println(msg)
	default:
		fmt.Println("No messages received")
	}

	msg := "Hello world"

	select {
	case messages <- msg:
		fmt.Println("Messages was sent")
	default:
		fmt.Println("Message was not sent")
	}
}

func closingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func(in <- chan int) {
		for {
			//more value indicates that programm closes the channel useing close()
			j, more := <- in
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("No more jobs")
				done <- true
				return
			}
		}
	}(jobs)

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("Sending job", i)
		time.Sleep(time.Second)
	}

	close(jobs)
	fmt.Println("Sent all jobs")

	<- done
}
