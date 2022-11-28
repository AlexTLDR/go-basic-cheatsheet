package main

import (
	d "counting_votes/datafile"
	"fmt"
	"log"
	"sort"
)

func main() {
	lines, err := d.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(lines)
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	//fmt.Println(counts)

	//Print the votes fo in a random order
	// for name, count := range counts {
	// 	fmt.Printf("Votes for %s: %d\n", name, count)
	// }

	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[name])
	}

}
