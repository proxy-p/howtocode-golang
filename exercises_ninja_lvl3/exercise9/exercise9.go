package main

import (
	"fmt"
)

func main() {
	favSport := "surfing"
	//using a SWITCH statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
