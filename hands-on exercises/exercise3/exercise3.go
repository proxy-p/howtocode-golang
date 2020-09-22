package main

import (
	"fmt"
)

//VARIABLES are package level scope, ASSIGNED with the following values

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
