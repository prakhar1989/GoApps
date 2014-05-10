package main

import "fmt"
import "strconv"
import "math"

type Stack struct {
	top  int
	data [10]int
}

func (s *Stack) Push(k int) {
	s.data[s.top] = k
	s.top++
}

func (s *Stack) Pop() int {
	s.top--
	return s.data[s.top]
}

func (s Stack) String() (str string) {
	for i := 0; i < s.top; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return
}

func varArgs(args ...int) {
	for _, n := range args {
		fmt.Println(n)
	}
}

func getMax(xs []int) (max int) {
	for _, n := range xs {
		if n > max {
			max = n
		}
	}
	return
}

func getMin(xs []int) (min int) {
	min = math.MaxInt8
	for _, n := range xs {
		if n < min {
			min = n
		}
	}
	return
}

func Map(xs []int, f func(int) int) []int {
    ys := make([]int, len(xs))
	for i, n := range xs {
		ys[i] = f(n)
	}
	return ys
}

func main() {
	xs := []int{44, 3, 2, 1, 10, 123}
	fmt.Println(getAverage(xs[:2]))
	fmt.Println(orderParams(7, 2))
	fmt.Println("Max", getMax([]int{10, 22, 41, 16}))
	fmt.Println("Min", getMin([]int{10, 22, 41, 16}))

	f := func(i int) int {
        return i * 5
	}
    fmt.Println(Map(xs, f))

}

func getAverage(xs []int) (avg float64) {
	var sum float64
	for _, v := range xs {
		sum += float64(v)
	}
	avg = sum / float64(len(xs))
	return
}

func orderParams(x int, y int) (lo int, hi int) {
	if x >= y {
		lo, hi = y, x
	} else {
		lo, hi = x, y
	}
	return
}
