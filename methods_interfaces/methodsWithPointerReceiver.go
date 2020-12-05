package main

import "fmt"

// declaring a new struct type
type car struct {
	brand string
	price int
}

// declaring a method for car type
// it changes the value it works on
func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
	// the changes are not seen to the outside world (pass by value)
}

// declaring a method with a pointer receiver
func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
	// the changes are seen the outside world

}

// method declarations are not permitted on named types that are themselves pointer types
type distance *int

// ERROR ->  invalid receiver type *distance (distance is a pointer type)
// func (d *distance) f() {
//  fmt.Println("say something")
// }

func main() {

	// declaring a car value
	myCar := car{brand: "Isuzu", price: 30000}

	// calling the method with a value receiver
	myCar.changeCar1("Opel", 21000)

	// no change due to the same pass by value mechanism  !!!
	fmt.Println(myCar) // => {Isuzu 30000}

	// calling the method with a pointer receiver. It changes the value!
	(&myCar).changeCar2("Subaru", 40000)
	fmt.Println(myCar) // -> {Subaru 40000}

	// declaring a pointer to car
	var yourCar *car
	yourCar = &myCar      // it stores the address of myCar
	fmt.Println(*yourCar) // -> {Subaru 40000}

	// calling the method on pointer type
	// valid ways to call the method:
	yourCar.changeCar2("Mazda", 20000)
	(*yourCar).changeCar2("Mazda", 20000)
	fmt.Println(*yourCar) // => {Mazda 20000}

	// in the above example both myCar and yourCar variables have been modified.
	fmt.Println(myCar) // => {Mazda 20000}

}
