package main

import (
	"fmt"
)

func main() {
	greeting := greet()
	fmt.Println(greeting)
}

type language string

// greet returns a greeting to the world.
func greet(l language) string {
	switch l {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
