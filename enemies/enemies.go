package main

import "fmt"

func main(){
	enemyHealth:= map[string]int{
		"Goblin": 50,
		"Orc": 100,
		"Dragon": 500,
	}
	count, exists := enemyHealth["Boss"]
	if exists {
		fmt.Println(count)
	} else {
		enemyHealth["Boss"] = 1000
	}
	enemyHealth["Goblin"] = 0
	delete(enemyHealth, "Goblin")

	for enemy, health := range enemyHealth {
    fmt.Printf("%s has %d health\n", enemy, health)
}

}
