package main

import "fmt"

func main() {
	var budget, price float32
	var gender, sport string
	var age int

	fmt.Scanln(&budget)
	fmt.Scanln(&gender)
	fmt.Scanln(&age)
	fmt.Scanln(&sport)

	if sport == "Gym" {
		if gender == "f" {
			price = 35
		} else if gender == "m" {
			price = 42
		}
	} else if sport == "Boxing" {
		if gender == "f" {
			price = 37
		} else if gender == "m" {
			price = 41
		}
	} else if sport == "Yoga" {
		if gender == "f" {
			price = 42
		} else if gender == "m" {
			price = 45
		}
	} else if sport == "Zumba" {
		if gender == "f" {
			price = 31
		} else if gender == "m" {
			price = 34
		}
	} else if sport == "Dances" {
		if gender == "f" {
			price = 53
		} else if gender == "m" {
			price = 51
		}
	} else if sport == "Pilates" {
		if gender == "f" {
			price = 37
		} else if gender == "m" {
			price = 39
		}
	}

	if age <= 19 {
		price *= 0.80
	}

	if price <= budget {
		fmt.Printf("You purchased a 1 month pass for %s.", sport)
	} else {
		fmt.Printf("You don't have enough money! You need $%.2f more.", price-budget)
	}
}
