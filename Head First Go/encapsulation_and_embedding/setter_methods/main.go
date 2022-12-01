package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{}
	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	date.SetMonth(12)
	date.SetDay(1)

	fmt.Println(date)
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) {
	d.Month = month
}

func (d *Date) SetDay(day int) {
	d.Day = day
}
