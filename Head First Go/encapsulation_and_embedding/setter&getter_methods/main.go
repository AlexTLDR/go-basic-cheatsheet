package main

import (
	"calendar/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	// Printing of the Getter Methods

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
