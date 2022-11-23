package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	year := now.Year()
	fmt.Println(year)
	brocken := "G# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(brocken)
	fmt.Println(fixed)
}
