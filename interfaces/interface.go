package main

import "fmt"

// Enemy interface requires an Attack() method
type Enemy interface {
	Attack()
}

// Orc struct
type Orc struct {
	Name string
}

// Orc implements Attack() with unique logic
func (o Orc) Attack() {
	fmt.Println(o.Name, "the Orc swings a club!")
}

// Dragon struct
type Dragon struct {
	Name string
}

// Dragon implements Attack() with unique logic
func (d Dragon) Attack() {
	fmt.Println(d.Name, "the Dragon breathes fire!")
}

func main() {
	// Create a slice of Enemy types
	enemies := []Enemy{
		Orc{Name: "Gorg"},
		Dragon{Name: "Smaug"},
	}

	// Iterate through the slice and call Attack() on each
	for _, enemy := range enemies {
		enemy.Attack()
	}
}
