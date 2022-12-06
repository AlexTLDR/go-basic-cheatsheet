package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

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
	scanDirectory("/home")

}
