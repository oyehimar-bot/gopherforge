package main

import "fmt"

func displayLevelUpMessage(newLevel int){
	fmt.Printf("Congratulations! You reached Level %d!\n", newLevel)
}

func main()  {
	displayLevelUpMessage(1)
}