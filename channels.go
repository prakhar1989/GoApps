package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(2 * time.Second)
	fmt.Println("done")

	// send a message to when done
	done <- true
}

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
	done := make(chan bool)
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go worker(done) // goroutine to work

	// block until we recieve a notification
	// from the worker on the channel
	<-done

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
