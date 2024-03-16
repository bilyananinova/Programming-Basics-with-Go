package main

import (
	"fmt"
	"strconv"
)

func main() {
	var capacity float64
	var suitcases int
	var input string

	fmt.Scanln(&capacity)
	fmt.Scanln(&input)

	for input != "End" {
		vol, _ := strconv.ParseFloat(input, 64)

		if (suitcases+1)%3 == 0 {
			vol *= 1.10
		}

		if capacity <= vol {
			fmt.Println("No more space!")
			break
		} else {
			capacity -= vol
			suitcases++
			fmt.Scanln(&input)
		}

	}

	if input == "End" {
		fmt.Println("Congratulations! All suitcases are loaded!")
	}

	fmt.Printf("Statistic: %d suitcases loaded.", suitcases)
}
