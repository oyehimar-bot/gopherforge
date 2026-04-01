package main

import "fmt"

// Stats Calculator
func calculateCharacterStats(baseStrength, baseAgility, baseIntelligence int) (totalAttackPower, totalMagicPower int) {
	totalAttackPower = baseStrength*2 + baseAgility/2
	totalMagicPower = baseIntelligence*3 + baseAgility/3
	return
}

// Item System
func useItem(itemName string, playerHealth int) (int, bool) {
	if itemName == "Health Potion" {
		playerHealth += 20
		return playerHealth, true
	} else if itemName == "Poison Vial" {
		playerHealth -= 15
		return playerHealth, true
	} else {
		fmt.Println("Unknown Item")
		return playerHealth, false
	}
}

func main() {
	attackPower, magicPower := calculateCharacterStats(10, 6, 8)
	fmt.Printf("Attack Power: %d\nMagic Power: %d\n", attackPower, magicPower)

	health, effect := useItem("Health Potion", 50)
	fmt.Println("Used Health Potion:", health, effect)
	health, effect = useItem("Poison Vial", 50)
	fmt.Println("Used Poison Vial:", health, effect)
	health, effect = useItem("Magic Scroll", 50)
	fmt.Println("Used Magic Scroll:", health, effect)
}
