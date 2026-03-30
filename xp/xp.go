package main

import "fmt"

func calculateExperienceGain(baseXP, levelMultiplier int) int {
	totalXp:= baseXP * levelMultiplier
	return totalXp
}

func main()  {
	totalXp:= calculateExperienceGain(50, 2)
	fmt.Printf("Total Experience Gained: %d\n", totalXp)
}