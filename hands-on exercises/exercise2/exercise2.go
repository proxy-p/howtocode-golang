package main

import (
	"fmt"
)

//Use keyword 'var' to DECLARE three VARIABLES. These variables have package level scope.
//VALUES are NOT ASSIGNED to the variables. Using IDENTIFIERS for the variables to make sure
//the variables are of the following TYPE

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
