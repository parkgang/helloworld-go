package main

import "fmt"

func sqyHi() {
	fmt.Println("Hi")
}

func main() {
	var myFunction func()
	myFunction = sqyHi
	myFunction()
}
