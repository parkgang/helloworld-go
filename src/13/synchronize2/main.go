package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "잔다")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "잠에서 깨다!")
}

func send(myChannel chan string) {
	reportNap("고루틴 전송", 2)
	fmt.Println("***전송 값***")
	myChannel <- "a"
	fmt.Println("***전송 값***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("고루틴 수신", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
