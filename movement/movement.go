package main

import "fmt"

func movePlayer(currentX, currentY int, direction string) (newX, newY int) {
	newX = currentX
	newY = currentY
	if direction == "north" {
		newY--
	} else if direction == "south" {
		newY++
	} else if direction == "east" {
		newX++
	} else if direction == "west" {
		newX--
	}
	return newX, newY
}

func MoveCharacter(direction string) error {
	switch direction {
	case "up", "down", "left", "right":
		fmt.Printf("Character moved %s\n", direction)
		return nil // nil means "no error"
	default:
		// fmt.Errorf creates a new error with a descriptive message
		return fmt.Errorf("invalid direction '%s'", direction)
	}
}

func main() {
	playerX := 10
	playerY := 15
	playerX, playerY = movePlayer(playerX, playerY, "north")
	fmt.Printf("Player's new position: X=%d, Y=%d\n", playerX, playerY)

	// 1. Test with a valid direction (Success Case)
	if err := MoveCharacter("up"); err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
	}

	// 2. Test with an invalid direction (Failure Case)
	// We capture the error in 'err' and immediately check it
	err := MoveCharacter("jump")
	if err != nil {
		// Print a custom warning (formatting the error value)
		fmt.Printf("⚠️ WARNING: %v\n", err)
	}
}
