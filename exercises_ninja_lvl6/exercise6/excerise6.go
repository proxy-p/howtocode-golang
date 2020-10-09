package main

import (
	"fmt"
)

func main() {

	//Building an ANONYMOUS FUNC
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}
