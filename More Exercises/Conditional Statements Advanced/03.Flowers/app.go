package main

import "fmt"

func main() {
	var chrysanthemums, roses, tulips int
	var season, holiday string
	var chrysanthemumsPrice, rosesPrice, tulipsPrice, total float32
	var isHodiday bool = false

	fmt.Scan(&chrysanthemums, &roses, &tulips, &season, &holiday)

	if holiday == "Y" {
		isHodiday = true
	}

	switch season {
	case "Spring", "Summer":
		chrysanthemumsPrice = 2.0
		rosesPrice = 4.10
		tulipsPrice = 2.50
	case "Autumn", "Winter":
		chrysanthemumsPrice = 3.75
		rosesPrice = 4.50
		tulipsPrice = 4.15
	}

	if isHodiday {
		chrysanthemumsPrice *= 1.15
		rosesPrice *= 1.15
		tulipsPrice *= 1.15
	}

	total = chrysanthemumsPrice*float32(chrysanthemums) + rosesPrice*float32(roses) + tulipsPrice*float32(tulips)

	if tulips > 7 && season == "Spring" {
		total *= 0.95
	}

	if roses >= 10 && season == "Winter" {
		total *= 0.90
	}

	if chrysanthemums+roses+tulips > 20 {
		total *= 0.80
	}

	total += 2

	fmt.Printf("%.2f", total)
}
