package main

import "fmt"
import "time"

func main() {

	timer1 := time.NewTimer(time.Second * 2) // expires in 2 seconds

	// blocks on timer's channel C until it sends a value indicating
	// that the timer expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
