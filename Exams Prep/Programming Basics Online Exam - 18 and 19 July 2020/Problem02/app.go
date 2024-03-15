package main

import "fmt"

func main() {
	var bagsPrice, bagsKg, price float32
	var days, count int

	fmt.Scan(&bagsPrice, &bagsKg, &days, &count)

	if bagsKg < 10 {
		price = bagsPrice * 0.20
	} else if bagsKg >= 10 && bagsKg <= 20 {
		price = bagsPrice * 0.50
	} else {
		price = bagsPrice
	}

	if days > 30 {
		price *= 1.10
	} else if days >= 7 && days <= 30 {
		price *= 1.15
	} else if days < 7 {
		price *= 1.40
	}

	fmt.Printf("The total price of bags is: %.2f lv.", price*float32(count))
}
