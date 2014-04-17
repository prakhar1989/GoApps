package main

import "github.com/stathat/go"
import "fmt"
import "time"
import "log"

func main() {
	code := "mH6wOpGb4bt7foFr"
	fmt.Println("Logging stathat")

	err := stathat.PostEZCount("jumps", code, 6)

	if err != nil {
		log.Printf("error posting ez count one: %v", err)
		return
	}
	ok := stathat.WaitUntilFinished(5 * time.Second)
	if ok {
		fmt.Println("ok")
	}
}
