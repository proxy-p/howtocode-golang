package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := bar(ii2)
	fmt.Println(n2)
}

//Creating a FUNC with an IDENTIFIER of 'FOO'
// Takes in a VARIADIC PARAMETER(..int) of TYPE INT
// Pass in a VALUE of TYPE []INT
func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	//RETURNS sum of all VALUES of TYPE INT passed in
	return total
}

//FUNC takes in a PARAMETER of TYPE []INT
func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	//RETURNS sum of all VALUES of TYPE INT passed in
	return total
}
