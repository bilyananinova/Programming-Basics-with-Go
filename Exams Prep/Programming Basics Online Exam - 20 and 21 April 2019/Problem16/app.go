package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var clients int
	var input string
	var totalMoney float64

	fmt.Scanln(&clients)

	for c := 1; c <= clients; c++ {
		scanner.Scan()
		input = scanner.Text()
		var purchaseCounter int
		var price float64

		for input != "Finish" {

			if input == "basket" {
				price += 1.50
			} else if input == "wreath" {
				price += 3.80
			} else if input == "chocolate bunny" {
				price += 7
			}

			purchaseCounter++

			scanner.Scan()
			input = scanner.Text()
		}

		if purchaseCounter%2 == 0 {
			price *= 0.80
		}

		totalMoney += price

		fmt.Printf("You purchased %d items for %.2f leva.\n", purchaseCounter, price)
	}

	average := totalMoney / float64(clients)
	fmt.Printf("Average bill per client is: %.2f leva.", average)
}
