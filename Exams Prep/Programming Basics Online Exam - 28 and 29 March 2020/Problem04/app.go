package main

import (
	"fmt"
	"math"
)

func main() {
	var days int
	var food, biscuits, eatenFoodByDogTotal, eatenFoodByCatTotal, eatenfoodByDogPerDay, eatenfoodByCatPerDay float64

	fmt.Scanln(&days)
	fmt.Scanln(&food)

	for d := 1; d <= days; d++ {
		fmt.Scanln(&eatenfoodByDogPerDay)
		fmt.Scanln(&eatenfoodByCatPerDay)

		eatenFoodByDogTotal += eatenfoodByDogPerDay
		eatenFoodByCatTotal += eatenfoodByCatPerDay

		if d%3 == 0 {
			biscuits += (eatenfoodByDogPerDay + eatenfoodByCatPerDay) * 0.10
		}

	}

	var totalEatenFood float64 = eatenFoodByDogTotal + eatenFoodByCatTotal

	fmt.Printf("Total eaten biscuits: %.0fgr.\n", math.Round(biscuits))

	percentEatenFromAllFood := totalEatenFood / food * 100
	fmt.Printf("%.2f%% of the food has been eaten.\n", percentEatenFromAllFood)

	percentEatenByDogFromEatenFood := eatenFoodByDogTotal / totalEatenFood * 100
	fmt.Printf("%.2f%% eaten from the dog.\n", percentEatenByDogFromEatenFood)

	percentEatenByCatFromEatenFood := eatenFoodByCatTotal / totalEatenFood * 100
	fmt.Printf("%.2f%% eaten from the cat.", percentEatenByCatFromEatenFood)
}
