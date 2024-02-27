package main

import "fmt"

func main() {
	var fuelType, haveDiscountCard string
	var qtyFuel, price float32

	fmt.Scan(&fuelType, &qtyFuel, &haveDiscountCard)

	var gasoline float32 = 2.22
	var diesel float32 = 2.33
	var gas float32 = 0.93

	if fuelType == "Gasoline" {
		if haveDiscountCard == "Yes" {
			gasoline -= 0.18
		}
		price = qtyFuel * gasoline
	} else if fuelType == "Diesel" {
		if haveDiscountCard == "Yes" {
			diesel -= 0.12
		}
		price = qtyFuel * diesel

	} else if fuelType == "Gas" {
		if haveDiscountCard == "Yes" {
			gas -= 0.08
		}
		price = qtyFuel * gas
	}

	if qtyFuel >= 20 && qtyFuel <= 25 {
		price *= 0.92
	} else if qtyFuel > 25 {
		price *= 0.90
	}

	fmt.Printf("%.2f lv.", price)
}
