package main

import (
	"fmt"
	"strings"
)

func main() {
	var days, wins, loses, totalwins, totalloses int
	var sport, result string
	var money, winMoney float32

	fmt.Scanln(&days)

	for d := 1; d <= days; d++ {
		fmt.Scanln(&sport)

		for sport != "Finish" {
			fmt.Scanln(&result)

			if result != "win" && result != "lose" {
				newSport := strings.Join([]string{sport, result}, " ")
				sport = newSport
				fmt.Scanln(&result)
			}

			if result == "win" {
				money += 20
				wins++
			} else {
				loses++
			}

			fmt.Scanln(&sport)
		}

		if wins > loses {
			money *= 1.10
		}

		totalloses += loses
		totalwins += wins
		winMoney += money
		money = 0
		wins = 0
		loses = 0
	}

	if totalwins > totalloses {
		winMoney *= 1.20
		fmt.Printf("You won the tournament! Total raised money: %.2f", winMoney)
	} else {
		fmt.Printf("You lost the tournament! Total raised money: %.2f", winMoney)
	}

}
