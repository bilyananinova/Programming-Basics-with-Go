package main

import "fmt"

func main() {
	var budget, fishermen int
	var season string
	var rent float32

	fmt.Scan(&budget, &season, &fishermen)

	if season == "Spring" {
		rent = 3000
	} else if season == "Summer" || season == "Autumn" {
		rent = 4200
	} else if season == "Winter" {
		rent = 2600
	}

	if fishermen <= 6 {
		rent *= 0.90
	} else if fishermen > 7 && fishermen <= 11 {
		rent *= 0.85
	} else if fishermen > 12 {
		rent *= 0.75
	}

	if fishermen%2 == 0 && season != "Autumn" {
		rent *= 0.95
	}

	if float32(budget) >= rent {
		left := float32(budget) - rent
		fmt.Printf("Yes! You have %.2f leva left.", left)
	} else {
		need := rent - float32(budget)
		fmt.Printf("Not enough money! You need %.2f leva.", need)
	}
}
