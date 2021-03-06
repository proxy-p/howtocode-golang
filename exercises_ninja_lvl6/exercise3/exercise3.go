package main

import (
	"fmt"
)

func main() {
	//Using DEFER to show a DEFFERED FUNC was ran AFTER the FUNC containing it exits
	defer foo()
	fmt.Println("Hello, playground")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	//This FUNC runs before the FIRST one
	fmt.Println("foo ran")
}
