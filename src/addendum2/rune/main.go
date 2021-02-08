package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "ABCDE"
	utf8String := "БГДЖИ"

	/// Example 1

	// 영어이므로 ASCII 집합으로 저장이 가능하여 총 "5byte"를 차지합니다.
	fmt.Println(len(asciiString))
	// 유니코드 문자 이므로 2~4byte으로 동적으로 변하여 총 "10byte"를 차지합니다.
	fmt.Println(len(utf8String))

	// len 함수는 바이트 단위로 반환하기 때문에 문자열의 길이를 문자 단위로 구하기 위해서는 아래의 함수를 사용합니다.
	fmt.Println(utf8.RuneCountInString(asciiString))
	fmt.Println(utf8.RuneCountInString(utf8String))

	/// Example 2

	asciiBytes := []byte(asciiString)
	utf8Bytes := []byte(utf8String)
	asciiBytesPartial := asciiBytes[3:]
	utf8BytesPartial := utf8Bytes[3:]
	// DE
	fmt.Println(string(asciiBytesPartial))
	// �ДЖИ => byte 단위로 슬라이스를 진행하여 1byte가 아닌 경우 2번째 문자가 반만 잘려 이상하게 출력되는 걸 확인할 수 있습니다.
	fmt.Println(string(utf8BytesPartial))

	asciiRunes := []rune(asciiString)
	utf8Runes := []rune(utf8String)
	asciiRunesPartial := asciiRunes[3:]
	utf8RunesPartial := utf8Runes[3:]
	// DE
	fmt.Println(string(asciiRunesPartial))
	// ЖИ => 바이트 슬라이스가 아닌 룬 슬라이스를 사용하면 문제없습니다.
	fmt.Println(string(utf8RunesPartial))

	/// Example 3

	// byte 단위로 출력하기 때문에 깨지는 문자열이 발생할 수 있습니다.
	for index, currentByte := range asciiBytes {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}
	for index, currentByte := range utf8Bytes {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}
	// 문자열을 순회할 때에는 바이트가 아닌 룬 단위로 순회하기 때문에 안전합니다!
	for index, currentByte := range asciiString {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}
	for index, currentByte := range utf8String {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}
}
