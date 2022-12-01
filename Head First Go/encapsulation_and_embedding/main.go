package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{Year: 2022, Month: 12, Day: 1}
	fmt.Println(date)
}

// Because users can enter invalid data, like Year: -1, Month: 13, Day: 45, in the encapsulating_and_embedding subfolder data validation will be used
