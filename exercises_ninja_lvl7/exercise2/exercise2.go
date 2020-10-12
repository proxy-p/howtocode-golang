package main

import (
	"fmt"
)

//Creating a TYPE of 'PERSON' with an underlying TYPE of STRUCT
type person struct {
	name string
}

func main() {
	//ASSIGNING a VALUE to TYPE 'PERSON'
	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1)
	//Calling FUNC 'CHANGEME'
	changeMe(&p1)
	fmt.Println(p1)
}

//Creating a FUNC and using the ADDRESS of TYPE 'PERSON' as a PARAMETER
//(PASSING in the ADDRESS of TYPE 'PERSON')
//This FUNC changes the VALUE stored in the ADDRESS for 'p1'
func changeMe(p *person) {
	p.name = "Miss Moneypenny"
	//This DEREFERENCES the ADDRESS of the VARIABLE
	// (*p).name = "Miss Moneyp"
}
