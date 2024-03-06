package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var bankBalance float64

	fmt.Scanln(&input)

	for input != "NoMoreMoney" {

		sum, _ := strconv.ParseFloat(input, 32)

		if sum < 0 {
			fmt.Println("Invalid operation!")
			break
		}

		fmt.Printf("Increase: %.2f\n", sum)
		bankBalance += sum

		fmt.Scanln(&input)
	}

	fmt.Printf("Total: %.2f", bankBalance)
}
