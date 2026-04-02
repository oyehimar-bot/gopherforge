package main

import "fmt"

func main()  {
	damage := []int{10, 20, 30, 40}
	damage = append(damage, 50, 60)
	damage= damage[1:]
	fmt.Println(damage)
}
