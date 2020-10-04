package main

import (
	"fmt"
)

func main() {
	//Creatings a SLICE type with various VALUES assigned to it
	x := []int{42, 43, 44, 88, 46, 47, 420, 49, 50, 51}
	for i, v := range x { //RANGE over the SLICE and PRINT the VALUES
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
