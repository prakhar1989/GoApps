package main

import "sort"
import "fmt"

type ByLength []string // alias for builtin string type

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 5, 10, 312}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits)) // using a custom function to sort
	fmt.Println(fruits)
}
