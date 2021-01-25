// Package datafile 파일로부터 샘플 데이터를 읽어 옵니다.
package datafile

import (
	"bufio"
	"log"
	"os"
)

func GetFloats(fileName string) ([3]float64, error) {
	var number [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

	}
}
