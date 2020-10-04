package main

import (
	"fmt"
)

func main() {
	x := "Money"
	//Showing and ELSE IF statement in action
	if x == "Money" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}
