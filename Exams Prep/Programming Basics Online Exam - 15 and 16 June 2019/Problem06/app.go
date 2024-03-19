package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input, bestMovie string
	var points, limit int

	scanner.Scan()
	input = scanner.Text()

	for input != "STOP" {
		var currentPoints int

		for i := 0; i < len(input); i++ {

			if input[i] >= 65 && input[i] <= 90 {
				currentPoints += int(input[i]) - len(input)
			}

			if input[i] >= 97 && input[i] <= 122 {
				currentPoints += int(input[i]) - (len(input) * 2)
			}

			if input[i] >= 48 && input[i] <= 57 {
				currentPoints += int(input[i])
			}

			space := string(int(input[i]))

			if space == " " {
				currentPoints += 32
			}
		}

		if points < currentPoints {
			points = currentPoints
			bestMovie = input
		}

		limit++

		if limit == 7 {
			fmt.Println("The limit is reached.")
			break
		}

		scanner.Scan()
		input = scanner.Text()
	}

	fmt.Printf("The best movie for you is %s with %d ASCII sum.", bestMovie, points)
}
