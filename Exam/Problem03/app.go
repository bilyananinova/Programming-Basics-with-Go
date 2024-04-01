package main

import "fmt"

func main() {
	var group int
	var season string
	var price float64

	fmt.Scanln(&group)
	fmt.Scanln(&season)

	switch season {
	case "spring":
		if group <= 5 {
			price = 50.0
		} else {
			price = 48.0
		}
	case "summer":
		if group <= 5 {
			price = 48.50
		} else {
			price = 45.0
		}
	case "autumn":
		if group <= 5 {
			price = 60.0
		} else {
			price = 49.50
		}
	case "winter":
		if group <= 5 {
			price = 86
		} else {
			price = 85
		}
	}

	total := price * float64(group)

	if season == "summer" {
		total *= 0.85
	}

	if season == "winter" {
		total *= 1.08
	}

	fmt.Printf("%.2f leva.", total)
}
