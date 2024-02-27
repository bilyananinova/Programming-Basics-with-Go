package main

import (
	"fmt"
	"math"
)

func main() {
	var x, z, workers int
	var y float64

	fmt.Scan(&x, &y, &z, &workers)

	totalGrapes := float64(x) * y
	var grapesForWine float64 = (totalGrapes * 0.40) / 2.5

	if grapesForWine >= float64(z) {
		fmt.Printf("Good harvest this year! Total wine: %.0f liters.\n", math.Floor(grapesForWine))
		left := math.Ceil(grapesForWine - float64(z))
		forWorker := math.Ceil(left / float64(workers))
		fmt.Printf("%.0f liters left -> %.0f liters per person.", left, forWorker)
	} else {
		needed := math.Floor(float64(z) - grapesForWine)
		fmt.Printf("It will be a tough winter! More %.0f liters wine needed.", needed)
	}
}
