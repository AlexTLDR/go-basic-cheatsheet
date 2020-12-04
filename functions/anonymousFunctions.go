package main

import "fmt"

// function that takes an int as an argument and returns another function that returns an int
func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	// declaring an anonymous functions
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!") // calling the anonymous function

	// calling the increment function. It returns an anonymous function
	a := increment(10)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
