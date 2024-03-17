package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var input, word string
	var maxScore int = math.MinInt32

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	for input != "End of words" {

		var end int = len(input)
		var score int = 0
		var isVow bool = false

		for i := 0; i < end; i++ {
			letter := input[i]

			score += int(letter)

			if i == 0 && (string(letter) == "A" || string(letter) == "a" || string(letter) == "E" || string(letter) == "e" ||
				string(letter) == "I" || string(letter) == "i" || string(letter) == "O" || string(letter) == "o" ||
				string(letter) == "U" || string(letter) == "u" || string(letter) == "Y" || string(letter) == "y") {
				isVow = true
			}

		}

		if !isVow {
			score = int(math.Floor(float64(score) / float64(end)))
		} else {
			score = score * end
		}

		if score >= maxScore {
			maxScore = score
			word = input
		}

		scanner.Scan()
		input = scanner.Text()

	}
	fmt.Printf("The most powerful word is %s - %d", word, maxScore)
}
