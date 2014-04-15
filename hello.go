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
	ages := make(map[string]int)
	ages["prakhar"] = 24
	ages["dhanam"] = 23
	ages["shanu"] = 20
	fmt.Println(ages)
	delete(ages, "shanu")
	fmt.Println(ages)
	_, not_present := ages["bob"] // not_present = false
	_, present := ages["prakhar"] // present = true
	fmt.Println(present, not_present)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	/** RANGES **/
	sum := 0
	for i, fr := range fruits { // first value is index
		fmt.Println("Adding fruit at index:", i)
		sum += fr
	}
	fmt.Println(sum)

	sum_ages := 0
	for _, v := range ages { // get key, value pairs
		sum_ages += v
	}
	fmt.Println(sum_ages)

	/** FUNCTIONS **/
	fmt.Println(factorial(5) == factRecurse(5)) // true
	fmt.Println(getSum(1, 2, 3, 4, 56, 12))     // 78
	// can take slices
	fmt.Println(getSum([]int{10, 10, 10}...)) // 30

	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
}

/** DEFINING FUNCTIONS **/
func factorial(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	return ans
}

// Returning multiple values
func vals() (int, int) {
	return 3, 4
}

// varying number of arguments
func getSum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// closures and anonymous functions
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// recursion
func factRecurse(n int) int {
	if n == 0 {
		return 1
	}
	return n * factRecurse(n-1)
}
