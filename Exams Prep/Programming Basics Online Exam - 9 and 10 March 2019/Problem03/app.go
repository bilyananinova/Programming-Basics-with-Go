package main

import "fmt"

func main() {
	var country, input string
	var difficulty, performance float64
	var maxScore float64 = 20

	fmt.Scanln(&country)
	fmt.Scanln(&input)

	if country == "Russia" {
		if input == "ribbon" {
			difficulty = 9.100
			performance = 9.400
		} else if input == "hoop" {
			difficulty = 9.300
			performance = 9.800
		} else if input == "rope" {
			difficulty = 9.600
			performance = 9.000
		}
	} else if country == "Bulgaria" {
		if input == "ribbon" {
			difficulty = 9.600
			performance = 9.400
		} else if input == "hoop" {
			difficulty = 9.550
			performance = 9.750
		} else if input == "rope" {
			difficulty = 9.500
			performance = 9.400
		}
	} else if country == "Italy" {
		if input == "ribbon" {
			difficulty = 9.200
			performance = 9.500
		} else if input == "hoop" {
			difficulty = 9.450
			performance = 9.350
		} else if input == "rope" {
			difficulty = 9.700
			performance = 9.150
		}
	}

	result := difficulty + performance

	fmt.Printf("The team of %s get %.3f on %s.\n", country, result, input)
	diff := ((maxScore - result) / maxScore) * 100
	fmt.Printf("%.2f%%", diff)
}
