package main

import "fmt"

func main() {
	var season, group, sport string
	var students, nights int
	var pricePerNight, price float32

	fmt.Scan(&season, &group, &students, &nights)

	switch season {
	case "Winter":
		if group == "girls" || group == "boys" {
			pricePerNight = 9.60

			if group == "girls" {
				sport = "Gymnastics"
			} else if group == "boys" {
				sport = "Judo"
			}

		} else if group == "mixed" {
			pricePerNight = 10
			sport = "Ski"
		}
	case "Spring":
		if group == "girls" || group == "boys" {
			pricePerNight = 7.20

			if group == "girls" {
				sport = "Athletics"
			} else if group == "boys" {
				sport = "Tennis"
			}

		} else if group == "mixed" {
			pricePerNight = 9.50
			sport = "Cycling"
		}
	case "Summer":
		if group == "girls" || group == "boys" {
			pricePerNight = 15

			if group == "girls" {
				sport = "Volleyball"
			} else if group == "boys" {
				sport = "Football"
			}

		} else if group == "mixed" {
			pricePerNight = 20
			sport = "Swimming"
		}
	}

	price = (pricePerNight * float32(nights)) * float32(students)

	if students >= 50 {
		price *= 0.50
	} else if students < 50 && students >= 20 {
		price *= 0.85
	} else if students < 20 && students >= 10 {
		price *= 0.95
	}

	fmt.Printf("%s %.2f lv.", sport, price)
}
