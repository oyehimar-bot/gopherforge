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
	actions := []string{
		"Player Jumping",
		"Enemy Shooting",
		"Background Music Playing",
	}

	for _, action := range actions {
		// Pass 'action' as parameter to avoid closure capture issue
		go func(act string) {
			playAction(act)
		}(action)
	}

	// Wait enough time for all goroutines to finish
	time.Sleep(100 * time.Millisecond)
}
