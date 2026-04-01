package main

import (
	"fmt"
	"strings"
)

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

// Party Introduction
func introduceParty(leaderName string, partyMembers ...string) {
	// if len(partyMembers) == 0{
	// 	fmt.Println("[LeaderName] has entered the dungeon alone.")
	// }
	if len(partyMembers) == 0 {
		// No party members
		// fmt.Printf("%s has entered the dungeon alone.\n", leaderName)
		fmt.Println("[LeaderName] has entered the dungeon alone.")
		return
	}

	var membersList string
	if len(partyMembers) == 1 {
		// Only one party member
		membersList = partyMembers[0]
	} else if len(partyMembers) == 2 {
		// Two party members, join with " and "
		membersList = partyMembers[0] + " and " + partyMembers[1]
	} else {
		// More than two party members
		// Join all but last with commas
		allButLast := strings.Join(partyMembers[:len(partyMembers)-1], ", ")
		// Add "and" before the last member
		membersList = allButLast + ", and " + partyMembers[len(partyMembers)-1]
	}

	fmt.Printf("%s has entered the dungeon with %s.\n", leaderName, membersList)
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

	introduceParty("Aria")
    introduceParty("Borin", "Cora")
    introduceParty("Dax", "Elena", "Finn")
    introduceParty("Gwen", "Hale", "Iris", "Jax")
}

