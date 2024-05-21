package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func transformStrings(s string, fn TransformFunc) string {
	return fn(s)
}

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func main() {
	fmt.Println(transformStrings("hello world!", Uppercase))
	fmt.Println(transformStrings("hello world!", Prefixer("Preixed with Prefixer func -> ")))
}
