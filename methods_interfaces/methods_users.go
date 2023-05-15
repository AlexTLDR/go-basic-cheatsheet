package main

import "fmt"

type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail implements a method with a pointer receiver
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// Values of type user can be used to call methods declared with a value receiver
	alex := user{"Alex", "alex@mail.com"}
	alex.notify()

	// Pointers of type user can also be used to call methods declared with a value receiver
	cami := &user{"Cami", "cami@email.com"}
	cami.notify()

	// Values of type user can be used to call methods declared with a pointer receiver
	alex.changeEmail("alex@gmail.com")
	alex.notify()

	// Pointers of type user can be used to call methods declared with a pointer receiver
	cami.changeEmail("cami@newemail.com")
	cami.notify()
}
