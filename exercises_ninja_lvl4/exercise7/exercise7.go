package main

import (
	"fmt"
)

func main() {
	//Creating 2 SLICES of TYPE STRING
	//PRINT out VALUES
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	//Storing (aggregating) the DATA into a MULTI-DIMENSIONAL SLICE
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	//RANGE over the RECORDS and PRINT the VALUES using an imbeded FOR loop
	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
