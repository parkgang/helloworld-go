package main

import (
	"fmt"
)

type MyType int

func (m *MyType) trans(message string) {
	fmt.Println(message)
	*m = 99
}

func main() {
	value := MyType(2)

	fmt.Println("value:", value)
	value.trans("변환 작업 중")
	fmt.Println("value:", value)
}
