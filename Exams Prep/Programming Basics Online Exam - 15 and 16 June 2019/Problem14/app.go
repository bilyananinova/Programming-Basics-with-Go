package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var budget, salary float64
	var input string
	fmt.Scanln(&budget)

	scanner.Scan()
	input = scanner.Text()
	// fmt.Scanln(&input)

	for input != "ACTION" {
		fmt.Scanln(&salary)

		if len(input) >= 15 {
			budget -= budget * 0.20
		} else {
			budget -= salary
		}

		if budget <= 0 {
			break
		}

		scanner.Scan()
		input = scanner.Text()
		// fmt.Scanln(&input)
	}

	if budget >= 0 {
		fmt.Printf("We are left with %.2f leva.", budget)
	} else {
		fmt.Printf("We need %.2f leva for our actors.", math.Abs(float64(budget)))
	}
}
