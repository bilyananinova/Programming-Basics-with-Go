package main

import "fmt"

func main() {

	var season string
	var kmPerMonth, kmPrice float32

	fmt.Scan(&season, &kmPerMonth)

	switch season {
	case "Spring", "Autumn":
		if kmPerMonth <= 5000 {
			kmPrice = 0.75
		} else if kmPerMonth > 5000 && kmPerMonth <= 10000 {
			kmPrice = 0.95
		} else if kmPerMonth > 10000 && kmPerMonth <= 20000 {
			kmPrice = 1.45
		}
	case "Summer":
		if kmPerMonth <= 5000 {
			kmPrice = 0.90
		} else if kmPerMonth > 5000 && kmPerMonth <= 10000 {
			kmPrice = 1.10
		} else if kmPerMonth > 10000 && kmPerMonth <= 20000 {
			kmPrice = 1.45
		}
	case "Winter":
		if kmPerMonth <= 5000 {
			kmPrice = 1.05
		} else if kmPerMonth > 5000 && kmPerMonth <= 10000 {
			kmPrice = 1.25
		} else if kmPerMonth > 10000 && kmPerMonth <= 20000 {
			kmPrice = 1.45
		}
	}

	total := ((kmPerMonth * kmPrice) * 0.90) * 4

	fmt.Printf("%.2f", total)
}
