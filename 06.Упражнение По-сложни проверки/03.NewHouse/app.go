package main

import "fmt"

func main() {
	var flowersType string
	var count, budget int
	var price, total float32

	fmt.Scan(&flowersType, &count, &budget)

	if flowersType == "Roses" {
		price = 5

		if count > 80 {
			price *= 0.90
		}
	} else if flowersType == "Dahlias" {
		price = 3.80

		if count > 90 {
			price *= 0.85
		}

	} else if flowersType == "Tulips" {
		price = 2.80

		if count > 80 {
			price *= 0.85
		}

	} else if flowersType == "Narcissus" {
		price = 3.0

		if count < 120 {
			price *= 1.15
		}

	} else if flowersType == "Gladiolus" {
		price = 2.50

		if count < 80 {
			price *= 1.20
		}

	}

	total = float32(count) * price

	if float32(budget) >= total {
		left := float32(budget) - total
		fmt.Printf("Hey, you have a great garden with %d %s and %.2f leva left.", count, flowersType, left)
	} else {
		need := total - float32(budget)
		fmt.Printf("Not enough money, you need %.2f leva more.", need)
	}
}
