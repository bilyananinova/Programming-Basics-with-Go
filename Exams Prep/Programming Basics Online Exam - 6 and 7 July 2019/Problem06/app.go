package main

import (
	"fmt"
	"math"
)

func main() {
	var playerName, player string
	var n int
	var bestScore int = math.MinInt32

	fmt.Scanln(&playerName)

	for playerName != "Stop" {
		var end int = len(playerName)
		var score int = 0

		for i := 0; i < end; i++ {
			fmt.Scanln(&n)
			letter := playerName[i]

			if n == int(letter) {
				score += 10
			} else {
				score += 2
			}
		}

		if score >= bestScore {
			bestScore = score
			player = playerName
		}

		fmt.Scanln(&playerName)
	}

	fmt.Printf("The winner is %s with %d points!", player, bestScore)
}
