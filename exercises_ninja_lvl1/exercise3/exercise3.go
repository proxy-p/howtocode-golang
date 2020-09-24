package main

import (
	"fmt"
)

//VARIABLES are package level scope, ASSIGNED with the following values

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//We are using 'fmt.Sprintf' to print all of the VALUES to one single string.
	//We ASSIGN the RETURNED value of TYPE string using the short declaration
	//operator to a VARIABLE with the IDENTIFIER “s”
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	//Then we PRINT 's'
}
