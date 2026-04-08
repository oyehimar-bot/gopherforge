package main

import "fmt"

type Weapon struct {
	Name       string
	Damage     int
	Durability float64
}

type Character struct {
	Name string
}

func Upgrade(w *Weapon) {
	w.Damage += 5
}

func BoostCharacter(c *Character) {
	if c == nil {
		fmt.Println("Character not found")
	} else {
		fmt.Println(c.Name)
	}
}

func main() {
	Gun := Weapon{
		Name:       "Gun",
		Damage:     20,
		Durability: 1.5,
	}

	w := &Gun
	fmt.Printf("%s damages is -%d with durability of %.1f before upgrade\n", w.Name, w.Damage, w.Durability)
	Upgrade(w)
	fmt.Printf("%s damages is -%d with durability of %.1f after upgrade\n", w.Name, w.Damage, w.Durability)

	Mykel := Character{Name: "Mykel"}
	var noName *Character = nil
	c := &Mykel
	BoostCharacter(c)
	BoostCharacter(noName)
}
