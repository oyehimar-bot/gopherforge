package main

import "fmt"

type Weapon struct{
	Name string
	Damage int
	Durabilty float64
}
func main()  {
	RustySword := Weapon{
		Name: "Usang",
		Damage: 20,
		Durabilty: 5.0,
	}
	fmt.Println(RustySword.Damage)
}