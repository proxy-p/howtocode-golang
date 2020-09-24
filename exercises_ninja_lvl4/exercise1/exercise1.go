package main

import (
	"fmt"
)

func main() {
	x := [5]int{0, 10, 22, 35, 48}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
