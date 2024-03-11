package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	var index int
	var bottles, pots, dishes int

	fmt.Scanln(&bottles)
	fmt.Scanln(&input)
	liquid := float32(bottles) * 750

	for input != "End" {
		container, _ := strconv.Atoi(input)
		index++

		if index == 3 {
			index = 0
			liquid -= float32(container) * 15
			pots += container
		} else {
			liquid -= float32(container) * 5
			dishes += container
		}

		if liquid < 0 {
			fmt.Printf("Not enough detergent, %.0f ml. more necessary!", math.Abs(float64(liquid)))
			break
		}

		fmt.Scanln(&input)
	}

	if input == "End" {
		fmt.Println("Detergent was enough!")
		fmt.Printf("%d dishes and %d pots were washed.\n", dishes, pots)
		fmt.Printf("Leftover detergent %.0f ml.\n", liquid)
	}
}
