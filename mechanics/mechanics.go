package main

import "fmt"

func main(){
	playerHealth := 80
	damage := 20
	armor := 5

	remainingHealth := playerHealth + armor - damage // Health after damage
	fmt.Printf("Remaining Health: %d\n", remainingHealth)

	criticalHealth:= playerHealth < 30 // When health is below 30
	fmt.Printf("Low HP: %t\n", criticalHealth)
}
