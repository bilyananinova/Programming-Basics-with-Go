package main

import (
	"fmt"
	"math"
)

func main() {
	var recordInSeconds, distance, time float64

	fmt.Scanln(&recordInSeconds)
	fmt.Scanln(&distance)
	fmt.Scanln(&time)

	totalTime := distance * time
	addTime := math.Floor((distance / 15)) * 12.5

	totalTime += addTime

	if totalTime < recordInSeconds {
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", totalTime)
	} else {
		diff := math.Abs(totalTime - recordInSeconds)
		fmt.Printf("No, he failed! He was %.2f seconds slower.", diff)
	}
}
