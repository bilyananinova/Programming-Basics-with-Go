package main

import (
	"fmt"
	"math"
)

func main() {
	var recordsInseconds, distanceInMeters, climbingTimePerMeter float64
	fmt.Scanln(&recordsInseconds)
	fmt.Scanln(&distanceInMeters)
	fmt.Scanln(&climbingTimePerMeter)

	totalTime := distanceInMeters * climbingTimePerMeter
	slowdown := math.Floor(float64(distanceInMeters/50)) * 30
	result := totalTime + slowdown

	if result < recordsInseconds {
		fmt.Printf("Yes! The new record is %.2f seconds.", result)
	} else {
		fmt.Printf("No! He was %.2f seconds slower.", result-recordsInseconds)
	}

}
