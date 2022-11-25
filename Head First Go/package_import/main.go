/*
In /home/<yourname>/go (or where is your GOROOT - /usr/lib/go/src/greeting ) create the src directory if it doesn't exist. Than create a greeting directory within the src directory. In the greeting folder,
create a greeting.go file and copy the below code in it.

package greeting

import "fmt"

func Hello() {
	fmt.Println("Hello!")
}

func Hi() {
	fmt.Println("Hi!")
}

Then create the below main.go file with the below code:
*/

package main

import "greeting"

func main() {
	greeting.Hello()
	greeting.Hi()
}
