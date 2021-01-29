package main

import (
	"fmt"
)

type Liters float64

func (re Liters) draw(value string) {
	fmt.Println("receiver:", re)
	fmt.Println("argument:", value)
}

func main() {
	busFuel := Liters(240.0)
	busFuel.draw("parameter input!")
}
