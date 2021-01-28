package main

import (
	"fmt"
)

type Liters float64
type Gallons float64

func (a Liters) draw(value string) {
	fmt.Println(value)
}

func main() {
	busFuel := Liters(240.0)
	busFuel.draw("parameter input!")
}

// notion 작성
// 다른언어는 this같은거 있지만 애는 리시버 매게변수로 하는거
// 리시버를 사용하는 메소드의 경우 그냥 호출이 안되는거
