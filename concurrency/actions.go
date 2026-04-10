package main

import (
	"fmt"
	"time"
)

func playAction(action string) {
	for i := 0; i < 5; i++ {
		fmt.Println(action)
		time.Sleep(10 * time.Millisecond)
	}
}
func main() {
	go playAction("Player Jumping")
	go playAction("Enemy shooting")
	go playAction("Background music playing")
	time.Sleep(10 * time.Millisecond)
}
