package main

import "fmt"

type Whistle string
type MakeSound string

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	//call methods on a value with the empty interface -> use type assertion
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}
}

func (w Whistle) Whistle() {
	fmt.Println(w)
}

func (w Whistle) MakeSound() {
	fmt.Println(w)
}

// call methods on a value with the empty interface -> use type assertion
func AcceptWhistle(whistle Whistle) {
	fmt.Println(whistle)
	whistle.MakeSound()
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Whistle("A whistle"))
}
