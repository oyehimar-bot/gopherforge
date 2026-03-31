package main

import "fmt"

const (
	//Enemies Stats
	MaxEnemiesOnScreen   = 10
	BossHealthMultiplier = 2.5
	GameVersion          = "1.0.0"
)

func main() {
	//Player Stats
	playerName := "Adventurer"
	playerHealth := 100
	playerMana := 50
	playerGold := 123.45
	isPoisoned := false

	// Dynamic Game Event
	newPlayerHealth := playerHealth + 25 // Player picks up a health potion
	newPlayerGold := playerGold + 10.75 // Enemy drops gold
	newIsPoisoned := true // Player enters poisoned area

	fmt.Println("Player:", playerName)
	fmt.Println("Original HP:", playerHealth)
	fmt.Println("Original MP:", playerMana)
	fmt.Println("Original GP:", playerGold)
	fmt.Println("Original poison state:", isPoisoned)
	fmt.Println("Numbers of enemies on screen:", MaxEnemiesOnScreen)
	fmt.Println("Boss HP is multiplied by", BossHealthMultiplier)
	fmt.Println("Game version:", GameVersion)
	fmt.Println("Updated HP:", newPlayerHealth)
	fmt.Println("Updated GP:", newPlayerGold)
	fmt.Println("Updated poison state:", newIsPoisoned)
}
