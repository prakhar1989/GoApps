package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// makes a channel of type string
	messages := make(chan string)

	// can take a max of 2 messages
	more_messages := make(chan string, 2)

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go func() {
		messages <- "ping" // sending a message in the channel
	}()

	// recieving the message
	msg := <-messages
	fmt.Println("Message recieved from channel:", msg)

	more_messages <- "buffered"
	more_messages <- "channel"

	fmt.Println(<-more_messages)
	fmt.Println(<-more_messages)

	var input string
	fmt.Print("Please enter something: ")
	fmt.Scanln(&input)
	fmt.Println("You entered:", input)
}
