package main

import (
	"fmt"
	"math"
)

func main() {
	var budget, clothesPrice float64
	var extrasCount int

	fmt.Scanln(&budget)
	fmt.Scanln(&extrasCount)
	fmt.Scanln(&clothesPrice)

	decor := budget * 0.10
	clothes := clothesPrice * float64(extrasCount)

	if extrasCount > 150 {
		clothes *= 0.90
	}

	total := decor + clothes
	diff := math.Abs(budget - total)

	if budget >= total {
		fmt.Println("Action!")
		fmt.Printf("Wingard starts filming with %.2f leva left.", diff)
	} else {
		fmt.Println("Not enough money!")
		fmt.Printf("Wingard needs %.2f leva more.", diff)
	}
}
