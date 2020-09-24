package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 88, 46, 47, 420, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
