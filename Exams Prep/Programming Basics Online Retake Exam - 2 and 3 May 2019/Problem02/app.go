package main

import "fmt"

func main() {
	var budget, fuel, costs float64
	var day string

	fmt.Scan(&budget, &fuel, &day)

	fuelCosts := fuel * 2.10
	costs = fuelCosts + 100

	if day == "Saturday" {
		costs *= 0.90
	} else if day == "Sunday" {
		costs *= 0.80
	}

	if budget >= costs {
		fmt.Printf("Safari time! Money left: %.2f lv.", budget-costs)
	} else {
		fmt.Printf("Not enough money! Money needed: %.2f lv.", costs-budget)
	}
}
