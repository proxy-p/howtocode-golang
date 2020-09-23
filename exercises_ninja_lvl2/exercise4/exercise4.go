package main

import (
	"fmt"
)

func main() {
	//ASSIGN a VALUE to a VARIBLE
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	//SHIFT the BITS of that INT over 1 position to the left, and ASSIGN that to a VARIABLE
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
