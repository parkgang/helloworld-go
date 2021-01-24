package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}

func main() {
	fmt.Printf("%.2f\n", paintNeeded(4.2, 3.0))
	fmt.Printf("%.2f\n", paintNeeded(5.2, 3.5))
	fmt.Printf("%.2f\n", paintNeeded(5.0, 3.3))
}
