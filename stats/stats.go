package main

import "fmt"

const (
	//Enemies Stats
	maxEnemiesOnScreen = 10
	bossHealthMultiplier = 2.5
	gameVersion = "1.0.0"
)
func main()  {
	//Player Stats
	playerName := "Adventurer"
	playerHealth := 100
	playMana := 50
	playerGold := 123.45
	isPoisoned := false

	fmt.Println("Player:", playerName)
	fmt.Println("HP:", playerHealth)
	fmt.Println("MP:", playMana)
	fmt.Println("GP:", playerGold)
	fmt.Println("Poisoned:", isPoisoned)
	fmt.Println("Numbers of enemies on screen:", maxEnemiesOnScreen)
	fmt.Println("Boss HP is multiplied by", bossHealthMultiplier)
	fmt.Println("Game version:", gameVersion)
}
