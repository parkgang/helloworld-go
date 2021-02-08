package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "ABCDE"
	utf8String := "БГДЖИ"

	// 영어이므로 ASCII 집합으로 저장이 가능하여 총 "5byte"를 차지합니다.
	fmt.Println(len(asciiString))
	// 유니코드 문자 이므로 2~4byte으로 동적으로 변하여 총 "10byte"를 차지합니다.
	fmt.Println(len(utf8String))

	// len 함수는 바이트 단위로 반환하기 때문에 문자열의 길이를 문자 단위로 구하기 위해서는 아래의 함수를 사용합니다.
	fmt.Println(utf8.RuneCountInString(asciiString))
	fmt.Println(utf8.RuneCountInString(utf8String))
}
