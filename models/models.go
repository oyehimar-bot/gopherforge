package main

import "fmt"

// Weapon character structure
type Weapon struct {
    Name       string
    Damage     int
    Durability float64
}

// Position struct for coordinates
type Position struct {
    X, Y float64
}

// Player struct with embedded Position and Health
type Player struct {
    Position Position // Corrected spelling
    Health   int
}

// Heal resets player's health to 100
func Heal(p *Player) {
    p.Health = 100
}

func main() {
    // Create a Weapon instance named "Rusty Sword"
    RustySword := Weapon{
        Name:       "Rusty Sword",
        Damage:     20,
        Durability: 5.0,
    }
    fmt.Println("Weapon Damage:", RustySword.Damage)

    // Create a Player with initial position
    Inyang := Player{
        Position: Position{
            X: 8.7,
            Y: 3.6,
        },
        Health: 50, // Example initial health
    }
    fmt.Println("Initial Player Y Position:", Inyang.Position.Y)

    // Move the player by updating coordinates
    Inyang.Position.X += 1.0
    Inyang.Position.Y += 2.0
    fmt.Println("Moved Player Position:", Inyang.Position)

    // Create a Player and reset health using Heal function
    p := Player{Health: 30}
    fmt.Println("Before Heal:", p.Health)
    Heal(&p)
    fmt.Println("After Heal:", p.Health)
}
