package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
	address
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

type staffAnonymous struct {
	name   string
	salary float64
	address
}

func main() {
	//create a postallAddress struct and use it to populate subscriber 1 and used anonymized the address field.
	// because of the anonymization, the field is no longer called as subscriber1.homeAddress = postalAddress
	postalAddress := address{street: "Test St", city: "Stuttgart", state: "BW", postalCode: "123456"}
	subscriber1 := defaultSubscriber("Alex")
	subscriber1.rate = 0.99
	subscriber1.address = postalAddress
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
	// setting the fields of the inner struct through the outer struct
	employee.homeAddress.street = "Stuifenstr"
	employee.homeAddress.city = "Ulm"
	employee.homeAddress.state = "BW"
	employee.homeAddress.postalCode = "6543210"
	fmt.Println("Employee name:", employee.name)
	fmt.Println("Salary:", employee.salary)
	// accessing the fields
	fmt.Println("Street:", employee.homeAddress.street)
	fmt.Println("City:", employee.homeAddress.city)
	fmt.Println("State:", employee.homeAddress.state)
	fmt.Println("Postal code:", employee.homeAddress.postalCode)
	// if I have used anonymization in the staff struct, the call would be like employee.address.street
	// even better, because of using anonymization, we can access the fields as embeded structs:

	// employee.homeAddress.street becomes employee.street

	var worker staffAnonymous
	worker.name = "John Doe Anonymous"
	worker.salary = 900000
	// setting the fields of the inner struct through the outer struct
	worker.street = "Stuifenstr Anonymous"
	worker.city = "Ulm Anonymous"
	worker.state = "BW Anonymous"
	worker.postalCode = "6543210 Anonymous"
	fmt.Println("Employee name:", worker.name)
	fmt.Println("Salary:", worker.salary)
	// accessing the fields
	fmt.Println("Street:", worker.street)
	fmt.Println("City:", worker.city)
	fmt.Println("State:", worker.state)
	fmt.Println("Postal code:", worker.postalCode)
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
	fmt.Println("Home Address:", s.address)
}

// Using pointers to modify a struct\
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
