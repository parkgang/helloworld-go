package main

import "fmt"

func main() {
	items := []string{"1", "2", "3"}
	items = append(items, "4", "5")
	fmt.Println(len(items))
	fmt.Printf("%#v\n", items)
}
