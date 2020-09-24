package main

import (
	"fmt"
)

type hotdog int

var x hotdog

//ASSIGING Variable 'y' as the UNDERLYING TYPE of your custom TYPE “x”
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	//using CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
