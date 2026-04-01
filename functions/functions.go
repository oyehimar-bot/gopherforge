package main

import "fmt"

func calculateCharacterStats(baseStrength, baseAgility, baseIntelligence int) (totalAttackPower, totalMagicPower int) {
	totalAttackPower = baseStrength * 2 + baseAgility / 2
	totalMagicPower = baseIntelligence * 3 + baseAgility / 3
	return
}

func main()  {
	attackPower, magicPower:= calculateCharacterStats(10, 6, 8)
	fmt.Printf("Attack Power: %d\nMagic Power: %d\n", attackPower, magicPower)
}
