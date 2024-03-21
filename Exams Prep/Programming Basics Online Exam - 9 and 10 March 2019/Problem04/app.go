package main

import (
	"fmt"
	"strconv"
)

func main() {
	var playerOneName, playerTwoName, input string
	var playerOneCard, playerTwoCard, playerOnePoints, playerTwoPoints int
	fmt.Scanln(&playerOneName)
	fmt.Scanln(&playerTwoName)
	fmt.Scanln(&input)

	for input != "End" {
		playerOneCard, _ = strconv.Atoi(input)
		fmt.Scanln(&playerTwoCard)

		if playerOneCard > playerTwoCard {
			playerOnePoints += playerOneCard - playerTwoCard
		} else if playerTwoCard > playerOneCard {
			playerTwoPoints += playerTwoCard - playerOneCard
		} else if playerOneCard == playerTwoCard {

			fmt.Println("Number wars!")

			fmt.Scanln(&input)
			playerOneCard, _ = strconv.Atoi(input)
			fmt.Scanln(&playerTwoCard)

			if playerOneCard > playerTwoCard {
				fmt.Printf("%s is winner with %d points", playerOneName, playerOnePoints)
			} else if playerTwoCard > playerOneCard {
				fmt.Printf("%s is winner with %d points", playerTwoName, playerTwoPoints)
			}

			break
		}

		fmt.Scanln(&input)
	}

	if input == "End" {
		fmt.Printf("%s has %d points\n", playerOneName, playerOnePoints)
		fmt.Printf("%s has %d points", playerTwoName, playerTwoPoints)
	}
}
