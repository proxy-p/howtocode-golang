package main

import (
	"fmt"
)

func main() {
	x := "Money"

	if x == "Money" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}
