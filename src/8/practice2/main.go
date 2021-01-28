package main

import (
	magazine "8/practice2/types"
	"fmt"
)

func main() {
	address := magazine.Address{
		Street:     "123 Oak st",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}
	subscriber := magazine.Subscriber{
		Name:    "Aman Singh",
		Address: address,
	}
	fmt.Println(subscriber)
}
