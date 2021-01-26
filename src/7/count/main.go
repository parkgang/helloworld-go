package main

import (
	"7/datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("../votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
