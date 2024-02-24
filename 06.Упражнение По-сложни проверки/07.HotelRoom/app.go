package main

import "fmt"

func main() {
	var month string
	var nights int
	var apartmentResult, studioResult float32

	fmt.Scan(&month, &nights)

	if month == "May" || month == "October" {
		studioResult = float32(nights) * 50.00
		apartmentResult = float32(nights) * 65.00

		if nights > 14 {
			studioResult *= 0.70
			apartmentResult *= 0.90
		} else if nights > 7 {
			studioResult *= 0.95
		}

	} else if month == "June" || month == "September" {
		studioResult = float32(nights) * 75.20
		apartmentResult = float32(nights) * 68.70

		if nights > 14 {
			studioResult *= 0.80
			apartmentResult *= 0.90
		}

	} else if month == "July" || month == "August" {
		studioResult = float32(nights) * 76.00
		apartmentResult = float32(nights) * 77.00

		if nights > 14 {
			apartmentResult *= 0.90
		}
	}

	fmt.Printf("Apartment: %.2f lv.\n", apartmentResult)
	fmt.Printf("Studio: %.2f lv.", studioResult)
}
