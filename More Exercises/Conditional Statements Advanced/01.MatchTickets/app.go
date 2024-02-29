package main

import "fmt"

func main() {
	var budget, ticketPrice float32
	var category string
	var group int

	fmt.Scan(&budget, &category, &group)

	if group >= 1 && group <= 4 {
		budget *= 0.25
	} else if group >= 5 && group <= 9 {
		budget *= 0.40
	} else if group >= 10 && group <= 24 {
		budget *= 0.50
	} else if group >= 25 && group <= 49 {
		budget *= 0.60
	} else if group >= 50 {
		budget *= 0.75
	}

	if category == "VIP" {
		ticketPrice = 499.99
	} else if category == "Normal" {
		ticketPrice = 249.99
	}

	price := float32(group) * ticketPrice

	if budget >= price {
		left := budget - price
		fmt.Printf("Yes! You have %.2f leva left.", left)
	} else {
		needed := price - budget
		fmt.Printf("Not enough money! You need %.2f leva.", needed)
	}

}
