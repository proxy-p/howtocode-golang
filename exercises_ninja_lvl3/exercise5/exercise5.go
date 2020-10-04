package main

import (
	"fmt"
)

func main() {
	//Print out the remainder (MODULUS) which is found for each number between 10 and 100 when it is divided by 4.
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}
