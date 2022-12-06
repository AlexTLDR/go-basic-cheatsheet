package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else { // if the panic value isn't an error, resume panicking with the same value.
		panic(p)
	}
}

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	// join the directory path and file name with a slash
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
			if err != nil {

			}
		} else {
			fmt.Println(filePath)
		}
	}

}

func main() {
	defer reportPanic()
	//panic("uncomment to test the else statement in the reportPanic function")
	scanDirectory("/")
}
