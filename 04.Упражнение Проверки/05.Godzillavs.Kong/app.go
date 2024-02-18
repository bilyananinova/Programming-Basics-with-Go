package main

import (
	"fmt"
	"math"
)

func main() {
	var budget, clothesPrice, decor, clothesSum float64
	var statistsCount int

	fmt.Scanln(&budget)
	fmt.Scanln(&statistsCount)
	fmt.Scanln(&clothesPrice)

	decor = budget * 0.10

	if statistsCount >= 150 {
		clothesPrice = clothesPrice * 0.90
	}

	clothesSum = float64(statistsCount) * clothesPrice
	diff := math.Abs(budget - clothesSum - decor)

	if budget >= clothesSum+decor {
		fmt.Printf("Action!\n")
		fmt.Printf("Wingard starts filming with %.2f leva left.", diff)
	} else {
		fmt.Printf("Not enough money!\n")
		fmt.Printf("Wingard needs %.2f leva more.", diff)
	}

}
