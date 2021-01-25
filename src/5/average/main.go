// Package main 프로그램은 숫자들의 평균을 계산합니다.
package main

import (
	"5/datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("../data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/sampleCount)
}
