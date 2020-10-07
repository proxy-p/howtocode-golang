package main

import (
	"fmt"
)

func main() {
	//Calling both FUNCs
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}

//Creating a FUNC that RETURNS an INT
func foo() int {
	return 42
}

//Creating a FUNC that RETURNS an INT & a STRING
func bar() (int, string) {
	return 1984, "Big Brother is Watching"
}
