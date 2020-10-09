package main

import (
	"fmt"
)

var x int

//ASSIGNING a VARIABLE as a TYPE FUNC
var g func()

func main() {

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//ASSINGING FUNC to a VARIABLE
	g = f
	//Calling the FUNC
	g()
	fmt.Printf("this is g %T\n", g)

	fmt.Println("done")
}
