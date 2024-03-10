package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	var allSteps int
	var target int = 10000

	for input != "GoingHome" {

		fmt.Scanln(&input)

		steps, _ := strconv.Atoi(input)

		allSteps += steps

		if allSteps >= target {
			fmt.Println("Goal reached! Good job!")
			fmt.Printf("%d steps over the goal!", allSteps-target)
			break
		}
	}

	if input == "GoingHome" {
		fmt.Scanln(&input)

		steps, _ := strconv.Atoi(input)

		allSteps += steps
		if allSteps >= target {
			fmt.Println("Goal reached! Good job!")
			fmt.Printf("%d steps over the goal!", allSteps-target)

		} else {
			fmt.Printf("%.0f more steps to reach goal.", math.Abs(float64(allSteps)-float64(target)))
		}
	}
}
