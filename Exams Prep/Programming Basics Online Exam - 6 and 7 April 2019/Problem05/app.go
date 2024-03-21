package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var filmsCount int
	var movie, maxMovie, minMovie string
	var rating, ratings float64

	var maxRating float64 = -10
	var minRating float64 = 10

	fmt.Scanln(&filmsCount)

	for i := 0; i < filmsCount; i++ {
		scanner.Scan()
		movie = scanner.Text()
		fmt.Scanln(&rating)

		if maxRating < rating {
			maxRating = rating
			maxMovie = movie
		}

		if minRating > rating {
			minRating = rating
			minMovie = movie
		}

		ratings += rating
	}

	average := ratings / float64(filmsCount)

	fmt.Printf("%s is with highest rating: %.1f\n", maxMovie, maxRating)
	fmt.Printf("%s is with lowest rating: %.1f\n", minMovie, minRating)
	fmt.Printf("Average rating: %.1f", average)
}
