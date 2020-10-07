package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
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

	//Creating a MAP of TYPE 'PERSON' with the key of LAST
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	//RANGE over MAP of TYPE 'PERSON'
	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		//RANGE over 'favFlavors'
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
