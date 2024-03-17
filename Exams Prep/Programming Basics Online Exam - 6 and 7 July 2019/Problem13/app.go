package main

import "fmt"

func main() {
	var town, typeOfTrip, vipDiscount string
	var days int
	var price, costs float32

	fmt.Scanln(&town)
	fmt.Scanln(&typeOfTrip)
	fmt.Scanln(&vipDiscount)
	fmt.Scanln(&days)

	switch town {
	case "Bansko", "Borovets":

		if typeOfTrip == "withEquipment" {
			price = 100

			if vipDiscount == "yes" {
				price *= 0.90
			}

		} else if typeOfTrip == "noEquipment" {
			price = 80

			if vipDiscount == "yes" {
				price *= 0.95
			}
		} else {
			fmt.Println("Invalid input!")
		}

	case "Varna", "Burgas":
		if typeOfTrip == "withBreakfast" {
			price = 130

			if vipDiscount == "yes" {
				price *= 0.88
			}
		} else if typeOfTrip == "noBreakfast" {
			price = 100

			if vipDiscount == "yes" {
				price *= 0.93
			}
		} else {
			fmt.Println("Invalid input!")
		}
	default:
		fmt.Println("Invalid input!")
	}

	if days > 7 {
		days--
	}

	costs = price * float32(days)

	if days <= 0 {
		fmt.Println("Days must be positive number!")
	}

	if costs > 0 {
		fmt.Printf("The price is %.2flv! Have a nice time!", costs)
	}
}
