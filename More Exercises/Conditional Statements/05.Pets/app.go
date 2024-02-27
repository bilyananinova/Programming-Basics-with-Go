package main

import (
	"fmt"
	"math"
)

func main() {
	var days, foodKgs int
	var dogFoodKgs, catFoodKgs, turtleFoodGrms float32

	fmt.Scan(&days, &foodKgs, &dogFoodKgs, &catFoodKgs, &turtleFoodGrms)

	turtleFoodKgs := turtleFoodGrms / 1000
	allNeededFood := dogFoodKgs*float32(days) + catFoodKgs*float32(days) + turtleFoodKgs*float32(days)

	if foodKgs >= int(allNeededFood) {
		fmt.Printf("%.0f kilos of food left.", math.Floor(float64(foodKgs)-float64(allNeededFood)))
	} else {
		fmt.Printf("%.0f more kilos of food are needed.", math.Ceil(float64(allNeededFood)-float64(foodKgs)))
	}
}
