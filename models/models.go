package main

import "fmt"

//Weapon character structure
type Weapon struct {
	Name      string
	Damage    int
	Durabilty float64
}

type Positon struct {
	X, Y float64
}

type Player struct {
	Positon Positon // A struct inside a struct
}

func main() {
	// Creating an instance using field names (recommended)
	RustySword := Weapon{
		Name:      "Usang",
		Damage:    20,
		Durabilty: 5.0,
	}
	fmt.Println(RustySword.Damage)

	Inyang := Player{
		Positon: Positon{ 
			X: 8.7,
			Y: 3.6,
		},
	}
	fmt.Println(Inyang.Positon.Y) // Access nested fields
}
