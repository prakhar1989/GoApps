package main

import (
	"fmt"
	"rand"
	"time"
)

type binFunc func(int, int) int

func main() {
	rand.Seed(time.Now().Unix())

	fns := []binFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
		func(x, y int) int { return x % y },
	}

	fn := fns[rand.Intn(len(fns))]
	fn := fns[2]

	x, y := 12, 5
	fmt.Println(fn(x, y))
}
