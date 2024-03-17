package main

import "fmt"

func main() {
	var drink, sugar string
	var order int
	var price, result float64

	fmt.Scanln(&drink)
	fmt.Scanln(&sugar)
	fmt.Scanln(&order)

	if drink == "Espresso" {
		if sugar == "Without" {
			price = 0.90
			price *= 0.65
		} else if sugar == "Normal" {
			price = 1
		} else if sugar == "Extra" {
			price = 1.20
		}

		if order >= 5 {
			price *= 0.75
		}
	} else if drink == "Cappuccino" {
		if sugar == "Without" {
			price = 1
			price *= 0.65
		} else if sugar == "Normal" {
			price = 1.20
		} else if sugar == "Extra" {
			price = 1.60
		}
	} else if drink == "Tea" {
		if sugar == "Without" {
			price = 0.50
			price *= 0.65
		} else if sugar == "Normal" {
			price = 0.60
		} else if sugar == "Extra" {
			price = 0.70
		}
	}

	result = price * float64(order)

	if result > 15 {
		result *= 0.80
	}

	fmt.Printf("You bought %d cups of %s for %.2f lv.", order, drink, result)
}
