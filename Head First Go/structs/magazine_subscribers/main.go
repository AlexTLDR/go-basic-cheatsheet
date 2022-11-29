package main

import "fmt"

type subscriber struct {
	name        string
	rate        float64
	active      bool
	homeAddress address
}

type staff struct {
	name        string
	salary      float64
	homeAddress address
}

type address struct {
	street     string
	city       string
	state      string
	postalCode string
}

func main() {
	postalAddress := address{street: "Test St", city: "Stuttgart", state: "BW", postalCode: "123456"}
	subscriber1 := defaultSubscriber("Alex")
	subscriber1.rate = 0.99
	subscriber1.homeAddress = postalAddress
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Cami")
	printInfo(subscriber2)

	subscriber3 := defaultSubscriber("Oreo")
	printInfo(subscriber3)
	applyDiscount(subscriber3)
	printInfo(subscriber3)

	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)

	var employee staff
	employee.name = "John Doe"
	employee.salary = 100000
	fmt.Println(employee.name)
	fmt.Println(employee.salary)
}

/* Using pointers for code optimization. Without pointers, functions receive a copy f the arguments they are called with.
That's why it's often a good idea to pass functions a pointer to a struct as structs are usually big values */

// func defaultSubscriber(name string) subscriber {
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	//return s
	return &s
}

// func printInfo(s subscriber) {
func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?:", s.active)
	fmt.Println("Home Address:", s.homeAddress)
}

// Using pointers to modify a struct\
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
