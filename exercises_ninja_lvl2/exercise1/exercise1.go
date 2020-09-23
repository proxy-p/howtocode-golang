package main

import (
	"fmt"
)

func main() {
	x := 142

	//PRINT a number in BINARY, DECIMAL, and HEX
	//Printf format %x reads arg #3, but call has 2 args
	//Needed to add the variable each time it was called
	fmt.Printf("%b\n%d\n%x", x, x, x)

}
