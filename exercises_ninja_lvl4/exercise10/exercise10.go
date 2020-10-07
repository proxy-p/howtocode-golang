package main

import (
	"fmt"
)

func main() {
	//Creating a MAP of TYPE STRING to store various VALUES of TYPE STRING
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)
	//APPENDING a record to the MAP with the following VALUES:
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	//DELETE a single record from MAP
	delete(m, `no_dr`)

	//RANGE over the MAP to PRINT each record along with their INDEX position in the SLICE
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
