package main

import "fmt"

const (
	//Enemies Stats
	maxEnemiesOnScreen   = 10
	bossHealthMultiplier = 2.5
	gameVersion          = "1.0.0"
)

func main() {
	//Player Stats
	playerName := "Adventurer"
	playerHealth := 100
	playMana := 50
	playerGold := 123.45
	isPoisoned := false

	// Dynamic Game Event
	newPlayerHealth := playerHealth + 25
	newPlayerGold := playerGold + 10.75
	newIsPoisoned := true

	fmt.Println("Player:", playerName)
	fmt.Println("Original HP:", playerHealth)
	fmt.Println("Original MP:", playMana)
	fmt.Println("Original GP:", playerGold)
	fmt.Println("Original poison state:", isPoisoned)
	fmt.Println("Numbers of enemies on screen:", maxEnemiesOnScreen)
	fmt.Println("Boss HP is multiplied by", bossHealthMultiplier)
	fmt.Println("Game version:", gameVersion)
	fmt.Println("Updated HP:", newPlayerHealth)
	fmt.Println("Updated GP:", newPlayerGold)
	fmt.Println("Updated poison state:", newIsPoisoned)
}
