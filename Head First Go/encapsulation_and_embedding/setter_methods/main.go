package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{}
	date.SetYear(2019)
	fmt.Println(date)
}

func (d Date) SetYear(year int) {
	d.Year = year
}
