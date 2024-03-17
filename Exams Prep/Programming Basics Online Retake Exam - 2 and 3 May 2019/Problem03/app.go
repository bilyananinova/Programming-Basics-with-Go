package main

import "fmt"

func main() {
	var duration, typeOfContract, addInternet string
	var months int
	var tax float64

	fmt.Scan(&duration, &typeOfContract, &addInternet, &months)

	if typeOfContract == "Small" {

		if duration == "one" {
			tax = 9.98
		} else if duration == "two" {
			tax = 8.58
		}

	} else if typeOfContract == "Middle" {

		if duration == "one" {
			tax = 18.99
		} else if duration == "two" {
			tax = 17.09
		}

	} else if typeOfContract == "Large" {

		if duration == "one" {
			tax = 25.98
		} else if duration == "two" {
			tax = 23.59
		}

	} else if typeOfContract == "ExtraLarge" {

		if duration == "one" {
			tax = 35.99
		} else if duration == "two" {
			tax = 31.79
		}

	}

	if addInternet == "yes" {
		if tax <= 10 {
			tax += 5.50
		} else if tax <= 30 {
			tax += 4.35
		} else if tax > 30 {
			tax += 3.85
		}
	}

	total := float64(months) * tax

	if duration == "two" {
		total = total * 0.9625
	}

	fmt.Printf("%.2f lv.", total)
}
