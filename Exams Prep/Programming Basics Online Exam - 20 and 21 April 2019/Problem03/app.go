package main

import "fmt"

func main() {
	var destination, dates string
	var nights int
	var price float64

	fmt.Scanln(&destination)
	fmt.Scanln(&dates)
	fmt.Scanln(&nights)

	if destination == "France" {
		if dates == "21-23" {
			price = 30
		} else if dates == "24-27" {
			price = 35
		} else if dates == "28-31" {
			price = 40
		}
	} else if destination == "Italy" {
		if dates == "21-23" {
			price = 28
		} else if dates == "24-27" {
			price = 32
		} else if dates == "28-31" {
			price = 39
		}
	} else if destination == "Germany" {
		if dates == "21-23" {
			price = 32
		} else if dates == "24-27" {
			price = 37
		} else if dates == "28-31" {
			price = 43
		}
	}

	total := price * float64(nights)

	fmt.Printf("Easter trip to %s : %.2f leva.", destination, total)
}
