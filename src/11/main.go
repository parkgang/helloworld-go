package main

import (
	"11/gadget"
	"11/mypkg"
	"fmt"
)

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	// player := gadget.TapePlayer{}
	// mixtape := []string{
	// 	"Jessie's Girl",
	// 	"Whip It",
	// 	"9 to 5",
	// }
	// playList(player, mixtape)

	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
