package main

import (
	"fmt"
)

func main() {
	var budget, price float64
	var season, accommodation, location string

	fmt.Scan(&budget, &season)

	if budget <= 1000 {
		accommodation = "Camp"

		if season == "Summer" {
			location = "Alaska"
			price = budget * 0.65
		} else if season == "Winter" {
			location = "Morocco"
			price = budget * 0.45
		}
	} else if budget >= 1000 && budget <= 3000 {
		accommodation = "Hut"

		if season == "Summer" {
			location = "Alaska"
			price = budget * 0.80
		} else if season == "Winter" {
			location = "Morocco"
			price = budget * 0.60
		}
	} else if budget >= 3000 {
		accommodation = "Hotel"

		if season == "Summer" {
			location = "Alaska"
		} else if season == "Winter" {
			location = "Morocco"
		}
		price = budget * 0.90
	}

	fmt.Printf("%s - %s - %.2f", location, accommodation, price)
}

