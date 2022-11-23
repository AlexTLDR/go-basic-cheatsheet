package main

import "fmt"

func main() {

	language := "golang"

	switch language {
	case "Rust": //values must be comparable (compare string to string)
		fmt.Println("You are learning Rust!")
		// an implicit break is added here
	case "Go", "golang": //compare language with "Go" OR "golang"
		fmt.Println("You are learning Go!")
	default:
		// the default clause the equivalent of the else clause of an if statement
		// and gets executed if no testing condition is true.
		fmt.Println("Any other programming language is a good start!")
	}

	n := 5
	// comparing the result of an expression which is bool to another bool value
	switch true {
	case n%2 == 0:
		fmt.Println("Even!")
	case n%2 != 0:
		fmt.Println("Odd!")
	default:
		fmt.Println("Never here!")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default
	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
