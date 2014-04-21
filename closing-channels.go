package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	fmt.Println("sent all jobs")
	close(jobs) // no more messges will be sent on jobs

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recieved job", j)
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return
			}
		}
	}()

	<-done // blocking until we recieve a notification from channel
}
