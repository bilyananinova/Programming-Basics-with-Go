package main

import (
	"fmt"
)

func main() {
	var tournamentName string
	var games, totalGames, wins, losts int

	fmt.Scanln(&tournamentName)

	for tournamentName != "End of tournaments" {
		fmt.Scanln(&games)
		totalGames += games

		var firstTeamResult int
		var secondTeamREsult int

		for i := 1; i <= games; i++ {
			fmt.Scanln(&firstTeamResult)
			fmt.Scanln(&secondTeamREsult)

			if firstTeamResult > secondTeamREsult {
				wins++
				fmt.Printf("Game %d of tournament %s: win with %d points.\n", i, tournamentName, firstTeamResult-secondTeamREsult)
			} else if secondTeamREsult > firstTeamResult {
				losts++
				fmt.Printf("Game %d of tournament %s: lost with %d points.\n", i, tournamentName, secondTeamREsult-firstTeamResult)
			}
		}

		fmt.Scanln(&tournamentName)
	}

	winPercent := float32(wins) / float32(totalGames) * 100
	fmt.Printf("%.2f%% matches win\n", winPercent)

	lostPercent := float32(losts) / float32(totalGames) * 100
	fmt.Printf("%.2f%% matches lost", lostPercent)
}
