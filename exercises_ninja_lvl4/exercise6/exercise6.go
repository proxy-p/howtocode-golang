package main

import (
	"fmt"
)

func main() {
	x := make([]string, 50, 500)
	//The BUILT-IN function MAKE creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T)
	//In this case, we are ALLOCATING an ARRAY with a CAPACITY of 500 STRINGS and creating a SLICE STRUCT with LENGTH of 50
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// put a VALUE into each of the 50 positions in the length of the SLICE
	//PRINT each POSITION and its VALUE
	for i := 0; i < 50; i++ {
		x[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// APPENDING more VALUES to the SLICE grows the length of the SLICE
	// but the underlying array "capacity" of 500 remains the same
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
