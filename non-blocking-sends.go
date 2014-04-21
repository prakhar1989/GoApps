package main

import "fmt"

func main() {
	messages := make(chan string)
	//signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	default:
		fmt.Println("no message recieved")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("message sent")
	default:
		fmt.Println("message not sent")
	}
}
