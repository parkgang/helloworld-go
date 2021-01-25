package main

import "fmt"

func main() {
	notes := [3]string{
		"1 인데",
		"2 라고",
		"3 신기해",
	}

	for index, item := range notes {
		fmt.Printf("%v: %v\n", index, item)
	}
}
