package main

import (
	d "counting_votes/datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := d.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
