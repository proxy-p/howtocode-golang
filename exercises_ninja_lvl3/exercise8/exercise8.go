package main

import (
	"fmt"
)

func main() {
	//Using a SWITCH statement with no switch expression specified
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}
