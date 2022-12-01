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

	// Embedding the Date type in an Event type

	event := calendar.Event{}
	err = event.SetYear(2023)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(26)
	if err != nil {
		log.Fatal(err)
	}
	// the getter methods for Date have been promoted to Event
	fmt.Println(event.Year())
	// long version -> get the Event's Date field, then call getter methods on it
	fmt.Println(event.Date.Month())
	fmt.Println(event.Day())

	// Encapsulating the Event Title field

	err = event.SetTitle("My b-day")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())

}
