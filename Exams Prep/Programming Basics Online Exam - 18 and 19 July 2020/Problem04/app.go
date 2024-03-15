package main

import (
	"fmt"
	"math"
)

func main() {
	var balls, points, redBalls, orangeBalls, yellowBalls, whiteBalls, otherBalls, divides int
	var color string

	fmt.Scanln(&balls)

	for ball := 0; ball < balls; ball++ {
		fmt.Scanln(&color)

		if color == "red" {
			points += 5
			redBalls++
		} else if color == "orange" {
			points += 10
			orangeBalls++
		} else if color == "yellow" {
			points += 15
			yellowBalls++
		} else if color == "white" {
			points += 20
			whiteBalls++
		} else if color == "black" {
			points = int(math.Floor(float64(points) / 2))
			divides++
		} else {
			otherBalls++
		}
	}

	fmt.Printf("Total points: %d\n", points)
	fmt.Printf("Red balls: %d\n", redBalls)
	fmt.Printf("Orange balls: %d\n", orangeBalls)
	fmt.Printf("Yellow balls: %d\n", yellowBalls)
	fmt.Printf("White balls: %d\n", whiteBalls)
	fmt.Printf("Other colors picked: %d\n", otherBalls)
	fmt.Printf("Divides from black balls: %d\n", divides)
}
