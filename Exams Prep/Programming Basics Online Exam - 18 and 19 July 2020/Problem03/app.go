package main

import "fmt"

func main() {
	var orderCount int
	var orderType, hasDelivery string
	var price, total float32

	fmt.Scanln(&orderCount)
	fmt.Scanln(&orderType)
	fmt.Scanln(&hasDelivery)

	if orderType == "90X130" {
		price = 110

		if orderCount > 60 {
			price = price * 0.92
		} else if orderCount > 30 {
			price = price * 0.95
		}
	} else if orderType == "100X150" {
		price = 140

		if orderCount > 80 {
			price = price * 0.90
		} else if orderCount > 40 {
			price = price * 0.94
		}
	} else if orderType == "130X180" {
		price = 190

		if orderCount > 50 {
			price = price * 0.88
		} else if orderCount > 20 {
			price = price * 0.93
		}
	} else if orderType == "200X300" {
		price = 250

		if orderCount > 50 {
			price = price * 0.86
		} else if orderCount > 25 {
			price = price * 0.91
		}
	}

	total = float32(orderCount) * price

	if hasDelivery == "With" {
		total += 60
	}

	if orderCount > 99 {
		total *= 0.96
	}

	if orderCount < 10 {
		fmt.Println("Invalid order")
	} else {
		fmt.Printf("%.2f BGN", total)
	}
}
