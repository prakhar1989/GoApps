package main

import "fmt"

// this function takes ONLY accepts a channel for
// sending values. This channel CANNOT recieve messages
func ping(pingChannel chan<- string, msg string) {
	pingChannel <- msg // send the message over the channel
}

// this function accepts a channel for recieving messages and another
// for recieving messages
func pong(pingChannel <-chan string, pongs chan<- string) {
	msg := <-pingChannel
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
