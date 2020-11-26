package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	a = int(b)

	fmt.Println(a, b)

	// var x int
	// x = "5"

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}
