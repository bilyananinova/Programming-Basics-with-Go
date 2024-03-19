package main

import (
	"fmt"
	"math"
)

func main() {
	var budget, price float32
	var serialsCount int
	var movie string

	fmt.Scanln(&budget)
	fmt.Scanln(&serialsCount)

	for i := 0; i < serialsCount; i++ {
		fmt.Scanln(&movie)
		fmt.Scanln(&price)

		if movie == "Thrones" {
			price *= 0.5
		} else if movie == "Lucifer" {
			price *= 0.6
		} else if movie == "Protector" {
			price *= 0.7
		} else if movie == "TotalDrama" {
			price *= 0.8
		} else if movie == "Area" {
			price *= 0.9
		}

		budget -= price

	}

	if budget >= 0 {
		fmt.Printf("You bought all the series and left with %.2f lv.", budget)
	} else {
		fmt.Printf("You need %.2f lv. more to buy the series!", math.Abs(float64(budget)))
	}
}
