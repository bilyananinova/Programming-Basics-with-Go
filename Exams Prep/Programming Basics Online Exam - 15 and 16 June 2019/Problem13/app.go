package main

import "fmt"

func main() {
	var budget, price float64
	var destination, season string
	var days int

	fmt.Scanln(&budget)
	fmt.Scanln(&destination)
	fmt.Scanln(&season)
	fmt.Scanln(&days)

	if destination == "Dubai" {
		if season == "Winter" {
			price = 45000
		} else if season == "Summer" {
			price = 40000
		}
	} else if destination == "Sofia" {
		if season == "Winter" {
			price = 17000
		} else if season == "Summer" {
			price = 12500
		}
	} else if destination == "London" {
		if season == "Winter" {
			price = 24000
		} else if season == "Summer" {
			price = 20250
		}
	}

	finalPrice := float64(days) * float64(price)

	if destination == "Dubai" {
		finalPrice *= 0.70
	}

	if destination == "Sofia" {
		finalPrice *= 1.25
	}

	if budget >= finalPrice {
		fmt.Printf("The budget for the movie is enough! We have %.2f leva left!", budget-finalPrice)
	} else {
		fmt.Printf("The director needs %.2f leva more!", finalPrice-budget)
	}
}
