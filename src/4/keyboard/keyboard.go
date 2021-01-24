// Package keyboard 키보드로부터 사용자의 입력을 읽어 옵니다.
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat 키보드로부터 부동 소수점 숫자를 읽어 옵니다.
// 이 함수는 읽은 숫자와 함께 에러 값을 반환합니다.
func GetFloat() (inputValue float64, _ error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
