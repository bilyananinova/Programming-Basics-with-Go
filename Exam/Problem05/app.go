package main

import "fmt"

func main() {
	var excursionSea, excursionMountain, profit int
	var input string

	fmt.Scanln(&excursionSea)
	fmt.Scanln(&excursionMountain)
	fmt.Scanln(&input)

	for input != "Stop" {
		if input == "sea" {
			if excursionSea > 0 {
				excursionSea--
				profit += 680
			}
		} else if input == "mountain" {
			if excursionMountain > 0 {
				excursionMountain--
				profit += 499
			}
		}

		if excursionSea == 0 && excursionMountain == 0 {
			fmt.Println("Good job! Everything is sold.")
			break
		}

		fmt.Scanln(&input)
	}

	fmt.Printf("Profit: %d leva.", profit)
}
