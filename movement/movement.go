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

func main(){
	playerX:= 10
	playerY:= 15
	playerX, playerY = movePlayer(playerX, playerY, "north")
	fmt.Printf("Player's new position: X=%d, Y=%d\n", playerX, playerY)
}