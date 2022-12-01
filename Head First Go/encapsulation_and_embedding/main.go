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
