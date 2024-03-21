package main

import (
	"fmt"
	"math"
)

func main() {
	var tournaments, startPoints, points, w int
	var result string

	fmt.Scanln(&tournaments)
	fmt.Scanln(&startPoints)

	for i := 0; i < tournaments; i++ {
		fmt.Scanln(&result)

		if result == "W" {
			points += 2000
			w++
		} else if result == "F" {
			points += 1200
		} else if result == "SF" {
			points += 720
		}

	}
	startPoints += points

	average := math.Floor(float64(points) / float64(tournaments))
	winPercent := float32(w) / float32(tournaments) * 100

	fmt.Printf("Final points: %d\n", startPoints)
	fmt.Printf("Average points: %.0f\n", average)
	fmt.Printf("%.2f%%", winPercent)
}
