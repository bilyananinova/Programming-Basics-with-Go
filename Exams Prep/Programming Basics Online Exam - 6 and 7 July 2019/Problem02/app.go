package main

import "fmt"

func main() {
	var budget, nightPrice, costs float32
	var nights, addCosts int

	fmt.Scanln(&budget)
	fmt.Scanln(&nights)
	fmt.Scanln(&nightPrice)
	fmt.Scanln(&addCosts)

	if nights > 7 {
		nightPrice *= 0.95
	}

	costs = float32(nights) * nightPrice

	costs += (budget * float32(addCosts)) / 100

	if budget >= costs {
		fmt.Printf("Ivanovi will be left with %.2f leva after vacation.", budget-costs)
	} else {
		fmt.Printf("%.2f leva needed.", costs-budget)
	}
}
