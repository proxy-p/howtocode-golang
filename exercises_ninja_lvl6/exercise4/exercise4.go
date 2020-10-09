package main

import (
	"fmt"
)

//Creating a TYPE 'PERSON'
type person struct {
	first string
	last  string
	age   int
}

//Creating a 'SPEAK' METHOD for a TYPE 'PERSON'
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	//ASSIGNING a VALUE of TYPE 'PERSON'
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	//Calling the 'SPEAK' METHOD of TYPE 'PERSON'
	p1.speak()
}
