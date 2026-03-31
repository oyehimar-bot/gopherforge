package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		heroName := os.Args[1]
		fmt.Printf("Greetings, noble %s! The elder awaits you.\n", heroName)
	} else {
		fmt.Println("Halt! State your name, traveler.")
	}
}
