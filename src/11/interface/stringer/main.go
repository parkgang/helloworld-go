package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%.2f L", l)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%.2f mL", m)
}

func main() {
	const value = 12.09248342
	fmt.Println(Gallons(value))
	fmt.Println(Liters(value))
	fmt.Println(Milliliters(value))
}
