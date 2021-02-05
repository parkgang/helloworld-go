package main

import (
	"14/prose"
	"fmt"
)

func main() {
	value := []string{
		"apple",
		"orange",
		"pear",
		"banana",
	}
	fmt.Println(prose.JoinWithCommas(value))
}
