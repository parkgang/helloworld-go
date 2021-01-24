package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat 키보드로 부터 사용자의 입력을 받아옵니다
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
