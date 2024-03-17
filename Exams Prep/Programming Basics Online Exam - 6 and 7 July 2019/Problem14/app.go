package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var h, w, percent int
	var input string
	fmt.Scanln(&h)
	fmt.Scanln(&w)
	fmt.Scanln(&percent)
	fmt.Scanln(&input)

	totalSqrt := h * w * 4
	wallsForPaint := math.Ceil(float64(totalSqrt) - float64(totalSqrt*percent/100))

	for input != "Tired!" {
		litters, _ := strconv.ParseFloat(input, 64)

		wallsForPaint -= litters

		if wallsForPaint <= 0 {
			if wallsForPaint < 0 {
				fmt.Printf("All walls are painted and you have %.0f l paint left!", math.Abs(wallsForPaint))
			} else if wallsForPaint == 0 {
				fmt.Println("All walls are painted! Great job, Pesho!")
			}
			break
		}
		fmt.Scanln(&input)
	}

	if input == "Tired!" {
		fmt.Printf("%.0f quadratic m left.", wallsForPaint)
	}
}
