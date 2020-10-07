package main

import (
	"fmt"
)

//Creating a TYPE of VEHICLE with an underlying TYPE of STRUCT
type vehicle struct {
	doors int
	color string
}

//Creating new TYPE of TRUCK & SEDAN with TYPE VEHICLE embeded
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	//Using COMPOSITE LITERAL to create a VALUE of TYPE TRUCK & SEDAN
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}
