package main

import (
	"fmt"
)

func main() {
	bd := 1975
	for {
		//Use an IF statement to PRINT all the years you have been alive
		if bd > 2020 {
			//if bd is greather than 2020, break the loop
			break
		}
		fmt.Println(bd)
		bd++
	}
}
