package main

import (
	"fmt"
)

func main() {
	//ASSIGN bd as an INT with the year you were born
	bd := 1975
	for bd <= 2020 {
		//Use FOR loop to PRINT bd and how many years you have been alive
		fmt.Println(bd)
		bd++
	}
}
