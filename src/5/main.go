package main

import "fmt"

func main() {
	notes := [3]string{
		"1 인데",
		"2 라고",
		"3 신기해",
	}

	for i := 0; i < 3; i++ {
		fmt.Println(notes[i])
	}
}
