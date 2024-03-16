package main

import "fmt"

func main() {
	var fruit, set string
	var order int
	var setPrice, totalCost float64

	fmt.Scanln(&fruit)
	fmt.Scanln(&set)
	fmt.Scanln(&order)

	if fruit == "Watermelon" {
		if set == "small" {
			setPrice = 2 * 56
		} else if set == "big" {
			setPrice = 5 * 28.70
		}
	} else if fruit == "Mango" {
		if set == "small" {
			setPrice = 2 * 36.66
		} else if set == "big" {
			setPrice = 5 * 19.60
		}
	} else if fruit == "Pineapple" {
		if set == "small" {
			setPrice = 2 * 42.10
		} else if set == "big" {
			setPrice = 5 * 24.80
		}
	} else if fruit == "Raspberry" {
		if set == "small" {
			setPrice = 2 * 20
		} else if set == "big" {
			setPrice = 5 * 15.20
		}
	}

	totalCost = setPrice * float64(order)

	if totalCost >= 400 && totalCost <= 1000 {
		totalCost *= 0.85
	} else if totalCost > 1000 {
		totalCost *= 0.50
	}

	fmt.Printf("%.2f lv.", totalCost)
}
