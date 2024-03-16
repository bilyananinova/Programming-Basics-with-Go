package main

import "fmt"

func main() {
	var walkMinutes, countWalks, caloriesPerDay int
	fmt.Scanln(&walkMinutes)
	fmt.Scanln(&countWalks)
	fmt.Scanln(&caloriesPerDay)

	totalMinutes := walkMinutes * countWalks
	totalCalories := totalMinutes * 5

	if totalCalories >= caloriesPerDay/2 {
		fmt.Printf("Yes, the walk for your cat is enough. Burned calories per day: %d.", totalCalories)
	} else {
		fmt.Printf("No, the walk for your cat is not enough. Burned calories per day: %d.", totalCalories)
	}
}
