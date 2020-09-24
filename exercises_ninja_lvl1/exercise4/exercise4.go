package main

import (
	"fmt"
)

//Creating our own TYPE as 'hotdog' and
type hotdog int

//Creating a VARIABLE of TYPE 'hotdog'
var x hotdog

func main() {
	//Print the VALUE and TYPE of VAR x
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
