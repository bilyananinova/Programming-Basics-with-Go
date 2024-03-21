package main

import (
	"fmt"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	var playerName, input string
	var currentPoints, successfulShots, unsuccessfulShots int
	var startPoints int = 301

	// scanner.Scan()
	// playerName = scanner.Text()
	fmt.Scanln(&playerName)
	fmt.Scanln(&input)

	for input != "Retire" {
		fmt.Scanln(&currentPoints)

		if input == "Double" {
			currentPoints *= 2
		} else if input == "Triple" {
			currentPoints *= 3
		}

		if (startPoints - currentPoints) >= 0 {
			startPoints -= currentPoints
			successfulShots++
		} else {
			unsuccessfulShots++
		}

		if startPoints == 0 {
			fmt.Printf("%s won the leg with %d shots.", playerName, successfulShots)
			break
		}

		fmt.Scanln(&input)
	}

	if input == "Retire" {
		fmt.Printf("%s retired after %d unsuccessful shots.", playerName, unsuccessfulShots)
	}
}
