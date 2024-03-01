package main

import "fmt"

func main() {
	var tournaments, startPoints, winPoints, winsCount int
	var stage string

	fmt.Scanln(&tournaments)
	fmt.Scanln(&startPoints)

	for i := 0; i < tournaments; i++ {

		fmt.Scanln(&stage)

		if stage == "W" {
			winPoints += 2000
			winsCount++
		} else if stage == "F" {
			winPoints += 1200
		} else if stage == "SF" {
			winPoints += 720
		}

	}

	average := winPoints / tournaments
	winsCountPercentage := float32(winsCount) / float32(tournaments) * 100

	fmt.Printf("Final points: %d\n", startPoints+winPoints)
	fmt.Printf("Average points: %d\n", average)
	fmt.Printf("%.2f%%", winsCountPercentage)

}
