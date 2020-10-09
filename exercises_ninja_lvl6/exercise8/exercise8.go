package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
}

//Building a FUNC that RETURNS a FUNC
func foo() func() int {
	return func() int {
		return 42
	}
}
