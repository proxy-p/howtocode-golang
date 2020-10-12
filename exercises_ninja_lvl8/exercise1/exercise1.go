package main

//Importing the necessary PACKAGE (ENCODING/JSON)
import (
	"encoding/json"
	"fmt"
)

//Creating a TYPE of VALUE 'USER' with an underlying TYPE of STRUCT
type user struct {
	First string
	Last  string
	Age   int
	//Creates a SLICE of TYPE STRING with a VALUE name of 'SAYINGS'
	Sayings []string
}

func main() {
	//ASSSIGNING VALUES to TYPE 'USER'
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	//ASSSIGNING VALUES to TYPE 'USER'
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	//ASSSIGNING VALUES to TYPE 'USER'
	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	//Using short hand operator to ASSIGN a NEW VARIABLE a NEW SLICE of []user
	users := []user{u1, u2, u3}

	fmt.Println(users)

	//Using short hand operator to ASSIGN 'bs' the FUNC of 'json.marshal()' which leverages the VALUE of
	//'users' to import the users slice
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
