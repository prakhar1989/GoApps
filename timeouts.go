package main

import "fmt"
import "time"

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// select proceeds with the first receive
	// thats ready, we timeout in this case
	select {
	case res := <-c1: // runs when received message from channel
		fmt.Println(res)
	case <-time.After(time.Second * 1): // runs after a second
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	// in this case, c2 wins
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	fmt.Println("I run at the end")
}
