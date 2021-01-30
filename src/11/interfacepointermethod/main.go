package main

import (
	"fmt"
)

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	var t Toggleable = &s // 포인터를 대신 할당합니다.
	t.toggle()
	t.toggle()
}
