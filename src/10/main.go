// setter, getter 연습
package main

import (
	"10/calendar"
	"fmt"
	"log"
)

func main() {
	var err error
	event := calendar.Event{}
	err = event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
