package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//Using APPEND to add VALUES to the end of the SLICE and PRINT
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	//Creating new SLICE of TYPE INT and APPENDING to first SLICE
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) //Need to UNFURL VALUES from Y SLICE into X SLICE
	fmt.Println(x)
}
