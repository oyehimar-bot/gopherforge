package main

import "fmt"

type Weapon struct {
	Name string
	Damage int
	Durability float64
}

func Upgrage(w *Weapon)  {
	w.Damage += 5
}

func main()  {
	Gun := Weapon{
		Name: "Gun",
		Damage: 20,
		Durability: 1.5,
	}

	w := &Gun
	fmt.Printf("%s damages is -%d with durability of %.1f\n", w.Name, w.Damage, w.Durability)
}