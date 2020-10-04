package main

import (
	"fmt"
)

func main() {
	//Using a COMPOSITE LITERAL to create an ARRAY which holds 5 VALUES of TYPE int
	//assign VALUES to each index position.
	x := [5]int{0, 10, 22, 35, 48} //This is the COMPOSITE LITERAL that creates the ARRAY with 5 VALUES
	for i, v := range x {          //RANGE over the ARRAY and PRINT the VALUES out.
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
