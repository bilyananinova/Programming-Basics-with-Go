package main

import "fmt"

func main() {
	var name, currentPlayer string
	var currentGoals, maxGoals int

	fmt.Scanln(&currentPlayer)

	for currentPlayer != "END" {
		fmt.Scanln(&currentGoals)

		if currentGoals > maxGoals {
			maxGoals = currentGoals
			name = currentPlayer
		}

		if currentGoals >= 10 {
			break
		}

		fmt.Scanln(&currentPlayer)

	}

	fmt.Printf("%s is the best player!\n", name)

	if maxGoals >= 3 {
		fmt.Printf("He has scored %d goals and made a hat-trick !!!", maxGoals)
	} else {
		fmt.Printf("He has scored %d goals.", maxGoals)
	}
}
