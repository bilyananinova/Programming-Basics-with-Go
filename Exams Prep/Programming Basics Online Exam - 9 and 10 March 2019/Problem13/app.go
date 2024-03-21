package main

import (
	"fmt"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	var stage, typeOfTicket, photo string
	var ticketsCount int
	var price float64

	// scanner.Scan()
	// stage = scanner.Text()

	fmt.Scanln(&stage)
	fmt.Scanln(&typeOfTicket)
	fmt.Scanln(&ticketsCount)
	fmt.Scanln(&photo)

	if stage == "Quarter final" {
		if typeOfTicket == "Standard" {
			price = 55.50
		} else if typeOfTicket == "Premium" {
			price = 105.20
		} else if typeOfTicket == "VIP" {
			price = 118.90
		}
	} else if stage == "Semi final" {
		if typeOfTicket == "Standard" {
			price = 75.88
		} else if typeOfTicket == "Premium" {
			price = 125.22
		} else if typeOfTicket == "VIP" {
			price = 300.40
		}
	} else if stage == "Final" {
		if typeOfTicket == "Standard" {
			price = 110.10
		} else if typeOfTicket == "Premium" {
			price = 160.66
		} else if typeOfTicket == "VIP" {
			price = 400
		}
	}

	result := price * float64(ticketsCount)

	if result > 4000 {
		result *= 0.75
	} else if result > 2500 {
		result *= 0.90
		if photo == "Y" {
			result += float64(ticketsCount) * 40
		}
	}

	fmt.Printf("%.2f", result)
}
