package main

import "fmt"

func main() {
	var temperature int

	var time, outfit, shoes string

	fmt.Scanln(&temperature)
	fmt.Scanln(&time)

	if time == "Morning" {

		if temperature >= 10 && temperature <= 18 {
			outfit = "Sweatshirt"
			shoes = "Sneakers"
		} else if temperature > 18 && temperature <= 24 {
			outfit = "Shirt"
			shoes = "Moccasins"
		} else if temperature >= 25 {
			outfit = "T-Shirt"
			shoes = "Sandals"
		}

	} else if time == "Afternoon" {

		if temperature >= 10 && temperature <= 18 {
			outfit = "Shirt"
			shoes = "Moccasins"
		} else if temperature > 18 && temperature <= 24 {
			outfit = "T-Shirt"
			shoes = "Sandals"
		} else if temperature >= 25 {
			outfit = "Swim Suit"
			shoes = "Barefoot"
		}

	} else if time == "Evening" {

		if temperature >= 10 && temperature <= 18 {
			outfit = "Shirt"
			shoes = "Moccasins"
		} else if temperature > 18 && temperature <= 24 {
			outfit = "Shirt"
			shoes = "Moccasins"
		} else if temperature >= 25 {
			outfit = "Shirt"
			shoes = "Moccasins"
		}

	}

	fmt.Printf("It's %d degrees, get your %s and %s.", temperature, outfit, shoes)
}
