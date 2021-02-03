package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "대기")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "가능!")
}

func send(myChannel chan string) {
	reportNap("전송", 2)
	fmt.Println("***전송 값***")
	myChannel <- "a"
	fmt.Println("***전송 값***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("수신", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
