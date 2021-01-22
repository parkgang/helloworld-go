package main

import (
	"fmt"
	"strings"
)

func Task1() {
	broken := "G# r#cks!"
	fixed := strings.NewReplacer("#", "o").Replace(broken)
	fmt.Println(fixed)
}
