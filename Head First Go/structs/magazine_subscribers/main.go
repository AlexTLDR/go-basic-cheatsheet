package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	subscriber1 := defaultSubscriber("Alex")
	subscriber1.rate = 0.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Cami")
	printInfo(subscriber2)

	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

// Using pointers to modify a struct\
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
