package main

import (
	"fmt"
)

type MyType string

func (m MyType) methed() {
	fmt.Println("리시버 변수:", m)
}

func (m *MyType) pointerMethod() {
	fmt.Println("리시버 변수 주소:", m)
}

func main() {
	value := MyType("안녕하세요")
	pointer := &value

	value.methed()
	value.pointerMethod()

	fmt.Println()

	pointer.methed()
	pointer.pointerMethod()
}
