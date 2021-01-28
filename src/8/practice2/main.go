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
		Name: "Aman Singh",
	}

	// 익명 필드를 사용하여 바로 접근
	subscriber.Street = address.Street
	subscriber.City = address.City
	subscriber.State = address.State
	subscriber.PostalCode = address.PostalCode

	fmt.Println(subscriber)
}
