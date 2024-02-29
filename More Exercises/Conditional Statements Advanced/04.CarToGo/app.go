package main

import "fmt"

func main() {
	var budget, price float32
	var season, carType, class string

	fmt.Scan(&budget, &season)

	if budget <= 100 {
		class = "Economy class"

		if season == "Summer" {
			carType = "Cabrio"
			price = budget * 0.35
		} else if season == "Winter" {
			carType = "Jeep"
			price = budget * 0.65
		}

	} else if budget > 100 && budget <= 500 {
		class = "Compact class"

		if season == "Summer" {
			carType = "Cabrio"
			price = budget * 0.45
		} else if season == "Winter" {
			carType = "Jeep"
			price = budget * 0.80
		}

	} else if budget > 500 {
		class = "Luxury class"
		carType = "Jeep"
		price = budget * 0.90
	}

	fmt.Println(class)
	fmt.Printf("%s - %.2f", carType, price)
}
