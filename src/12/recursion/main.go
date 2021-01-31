package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	const rootPath = "/Users/kyungeun/Sources/repos/ruddms936/helloworld-go/src/12/recursion"
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
