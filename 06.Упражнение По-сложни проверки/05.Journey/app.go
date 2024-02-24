package main

import "fmt"

func main() {
	var budget, costs float32
	var season, destination, restType string

	fmt.Scan(&budget, &season)

	if budget <= 100 {
		destination = "Bulgaria"

		if season == "summer" {
			costs = budget * 30 / 100
			restType = "Camp"
		} else if season == "winter" {
			costs = budget * 70 / 100
			restType = "Hotel"
		}

	} else if budget >= 100 && budget <= 1000 {
		destination = "Balkans"

		if season == "summer" {
			costs = budget * 40 / 100
			restType = "Camp"
		} else if season == "winter" {
			costs = budget * 80 / 100
			restType = "Hotel"
		}

	} else if budget >= 1000 {
		destination = "Europe"
		costs = budget * 90 / 100
		restType = "Hotel"
	}

	fmt.Printf("Somewhere in %s\n", destination)
	fmt.Printf("%s - %.2f", restType, costs)

}
