package main

import (
	"fmt"
)

//Creating a TYPE which will have an underlying type of STRUCT
type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	//Creating 	two VALUES of TYPE 'PERSON'
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	//Printing TYPE 'PERSON' VALUES
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	//RANGE over 'PERSON' favFlavors
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}
