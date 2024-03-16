package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var totalFoodInKg int
	var input string
	fmt.Scanln(&totalFoodInKg)
	fmt.Scanln(&input)

	totalFoodInGr := totalFoodInKg * 1000

	for input != "Adopted" {
		eatenFoodInGr, _ := strconv.Atoi(input)

		totalFoodInGr -= eatenFoodInGr

		fmt.Scanln(&input)
	}

	if totalFoodInGr < 0 {
		neededFood := math.Abs(float64(totalFoodInGr))
		fmt.Printf("Food is not enough. You need %.0f grams more.", neededFood)
	} else {
		fmt.Printf("Food is enough! Leftovers: %d grams.", totalFoodInGr)
	}
}
