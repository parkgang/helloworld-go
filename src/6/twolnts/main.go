package main

import "fmt"

func severalInts(a int, b bool, numbers ...string) {
	fmt.Println(a, b, numbers)
}

func main() {
	severalInts(1, true)
	severalInts(2, false)
	severalInts(2, false, "박", "경", "은")
}
