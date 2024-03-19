package main

import (
	"fmt"
	"math"
)

func main() {
	var filmmingTime, scenesCount, sceneTime int
	fmt.Scanln(&filmmingTime)
	fmt.Scanln(&scenesCount)
	fmt.Scanln(&sceneTime)

	prepTime := float64(filmmingTime) * 0.15
	neededTime := float64(scenesCount*sceneTime) + prepTime

	if float64(filmmingTime) >= neededTime {
		leftTime := float64(filmmingTime) - neededTime
		fmt.Printf("You managed to finish the movie on time! You have %.0f minutes left!", math.Round(leftTime))
	} else {
		need := neededTime - float64(filmmingTime)
		fmt.Printf("Time is up! To complete the movie you need %.0f minutes.", need)
	}
}
