package main

import "fmt"

type Whistle string

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
}

func (w Whistle) Whistle() {
	fmt.Println(w)
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Whistle("A whistle"))
}
