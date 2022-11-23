package main

import "fmt"

func main() {

	// each key corresponds to an index of the array
	grades := [3]int{ //the keyed elements can be in any order
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) // -> [5 10 7]

	accounts := [3]int{
		2: 50,
	}
	fmt.Println(accounts) //[0 0 50]

	names := [...]string{
		4: "Jan",
	}
	fmt.Println(len(names)) // -> 5

	// un unkeyed element gets its index from the last keyed element
	cities := [...]string{
		5:        "Munchen",
		"Berlin", // this is at index 6
		1:        "Hamburg",
	}
	fmt.Printf("%#v\n", cities) // -> [7]string{"", "Hamburg", "", "", "", "Munchen", "Berlin"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // => [false false false false false true true]
}
