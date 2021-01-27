package main

import "fmt"

func main() {
	items := map[string]string{
		"이름": "박경은",
		"나이": "22",
	}
	name, ok := items["이름"]
	fmt.Println(name, ok)

	age, ok := items["나이"]
	fmt.Println(age, ok)

	sex, ok := items["성"]
	fmt.Println(sex, ok)
}
