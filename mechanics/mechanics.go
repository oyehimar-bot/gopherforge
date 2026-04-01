package main

import "fmt"

func main() {
	//Damage Calculation
	playerHealth := 80
	damage := 20
	armor := 5

	remainingHealth := playerHealth + armor - damage // Health after damage
	fmt.Printf("Remaining Health: %d\n", remainingHealth)

	criticalHealth := playerHealth < 30 // When health is below 30
	fmt.Printf("Low HP: %t\n", criticalHealth)

	//Shop Logic
	goldCollected := 120
	itemCost := 75
	inventorySlots := 10
	freeSlots := inventorySlots - 2 // Assuming 2 slots are used

	canAfford := goldCollected >= itemCost
	hasSpace := freeSlots <= 10
	canBuy := canAfford && hasSpace
	fmt.Printf("Affordability is %t\n", canBuy)

	//Level Up Condition
	XP := 850
	questXP := 200
	const levelUpXP = 1000

	XP += questXP
	levelUp := XP >= levelUpXP
	fmt.Printf("Leveling Up is %t\n", levelUp)
}
