package main

import (
	"fmt"
	"strings"
)

func main() {

	// an anonymous struct is a struct with no explicitly defined struct type alias.
	jan := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Jan",
		lastName:  "Schwartz",
		age:       30,
	}

	fmt.Printf("%#v\n", jan)
	// =>struct { firstName string; lastName string; age int }{firstName:"Jan", lastName:"Schwartz", age:30

	//** ANONYMOUS FIELDS **//

	// fields type becomes fields name.
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"The Hunger Games", 10.2, false}
	fmt.Printf("%#v\n", b1) // => main.Book{string:"The Hunger Games", float64:10.2, bool:false}

	fmt.Println(b1.string) // => The Hunger Games

	// mixing anonymous with named fields:
	type Employee1 struct {
		name   string
		salary int
		bool
	}

	e := Employee1{"John", 40000, false}
	fmt.Printf("%#v\n", e) // => main.Employee1{name:"John", salary:40000, bool:false}
	e.bool = true          // changing a field

	fmt.Println(strings.Repeat("#", 10))

	//** EMBEDDED STRUCTS **//
	// An embedded struct is just a struct that acts like a field inside another struct.

	// define a new struct type
	type Contact struct {
		email, address string
		phone          string
	}

	// define a struct type that contains another struct as a field
	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	// declaring a value of type Employee
	john := Employee{
		name:   "John Doe",
		salary: 5000,
		contactInfo: Contact{
			email:   "jdoe@company.com",
			address: "Str 2, New York",
			phone:   "042324234",
		},
	}

	fmt.Printf("%+v\n", john)
	// => {name:John Doe salary:5000 contactInfo:{email:jdoe@company.com address:Str 2, New York phone:295619381404}}

	// accessing a field
	fmt.Printf("Employee's salary: %d\n", john.salary)

	// accessing a field from the embedded struct
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // => Employee's email:jdoe@company.com

	// updating a field
	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email)
	// =>  Employee's new email address:new_email@thecompany.com
}
