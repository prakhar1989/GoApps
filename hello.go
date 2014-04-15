package main

import "fmt"

func main() {
	const name string = "Prakhar"
	fmt.Println("4+4 =", 4+4)
	fmt.Println("Hello, ", name)

	/** FOR LOOPS **/
	for j := 2; j <= 9; j++ {
		fmt.Print(j, " ")
	}

	fmt.Println()

	/** IF ELSE **/
	apples := 10
	if apples > 5 {
		fmt.Println("Lots of apples!")
	} else {
		fmt.Println("Too few apples")
	}

	/** ARRAYS **/
	// standard way
	var friends [3]string
	friends[0] = "alice"
	friends[1] = "bob"
	friends[2] = "zoe"
	fmt.Println(friends)

	// another way of declaring arrays
	fruits := [5]int{4, 2, 10, 23, 20}
	fmt.Println("Total number of fruits", len(fruits))

	/** SLICES **/
	s := make([]string, 2)
	s[0] = "prakhar"
	s = append(s, "dhanam") // python like
	fmt.Println(s)

	/** MAPS **/
	friend_ages := make(map[string]int)
	friend_ages["prakhar"] = 24
	friend_ages["dhanam"] = 23
	friend_ages["shanu"] = 20
	fmt.Println(friend_ages)
}
