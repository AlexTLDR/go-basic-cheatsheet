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
}
