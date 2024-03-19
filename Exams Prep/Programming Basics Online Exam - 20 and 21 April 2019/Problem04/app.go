package main

import "fmt"

func main() {
	var playerOneEggs, playerTwoEggs int
	var input string

	fmt.Scanln(&playerOneEggs)
	fmt.Scanln(&playerTwoEggs)
	fmt.Scanln(&input)

	for input != "End" {

		if input == "one" {
			playerTwoEggs--
		} else if input == "two" {
			playerOneEggs--
		}

		if playerOneEggs == 0 || playerTwoEggs == 0 {
			break
		}

		fmt.Scanln(&input)
	}

	if playerOneEggs == 0 {
		fmt.Printf("Player one is out of eggs. Player two has %d eggs left.", playerTwoEggs)
	} else if playerTwoEggs == 0 {
		fmt.Printf("Player two is out of eggs. Player one has %d eggs left.", playerOneEggs)
	} else {
		fmt.Printf("Player one has %d eggs left.\n", playerOneEggs)
		fmt.Printf("Player two has %d eggs left.", playerTwoEggs)
	}
}
