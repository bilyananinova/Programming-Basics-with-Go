package main

import "fmt"

func main() {
	var product string
	var town string
	var qty float32

	fmt.Scanln(&product)
	fmt.Scanln(&town)
	fmt.Scanln(&qty)

	switch product {
	case "coffee":
		if town == "Sofia" {
			fmt.Println(qty * 0.50)
		} else if town == "Plovdiv" {
			fmt.Println(qty * 0.40)
		} else if town == "Varna" {
			fmt.Println(qty * 0.45)
		}
	case "water":
		if town == "Sofia" {
			fmt.Println(qty * 0.80)
		} else if town == "Plovdiv" {
			fmt.Println(qty * 0.70)
		} else if town == "Varna" {
			fmt.Println(qty * 0.70)
		}
	case "beer":
		if town == "Sofia" {
			fmt.Println(qty * 1.20)
		} else if town == "Plovdiv" {
			fmt.Println(qty * 1.15)
		} else if town == "Varna" {
			fmt.Println(qty * 1.10)
		}
	case "sweets":
		if town == "Sofia" {
			fmt.Println(qty * 1.45)
		} else if town == "Plovdiv" {
			fmt.Println(qty * 1.30)
		} else if town == "Varna" {
			fmt.Println(qty * 1.35)
		}
	case "peanuts":
		if town == "Sofia" {
			fmt.Println(qty * 1.60)
		} else if town == "Plovdiv" {
			fmt.Println(qty * 1.50)
		} else if town == "Varna" {
			fmt.Println(qty * 1.55)
		}
	}
}
