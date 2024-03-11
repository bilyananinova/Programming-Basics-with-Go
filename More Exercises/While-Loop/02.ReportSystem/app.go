package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var index, sum, countCS, countCC int
	var averageCS, averageCC, cashSum, cardSum float32

	fmt.Scanln(&input)

	charitySum, _ := strconv.Atoi(input)

	for sum <= charitySum {
		fmt.Scanln(&input)

		if input == "End" {
			break
		}

		transaction, _ := strconv.Atoi(input)

		if index%2 == 0 && input != "End" {
			if transaction > 100 {
				fmt.Println("Error in transaction!")
			} else {
				fmt.Println("Product sold!")
				sum += transaction
				cashSum += float32(transaction)
				countCS++
			}
		} else if index%2 == 1 && input != "End" {
			if transaction < 10 {
				fmt.Println("Error in transaction!")
			} else {
				fmt.Println("Product sold!")
				sum += transaction
				cardSum += float32(transaction)
				countCC++
			}
		}

		index++

		if sum >= charitySum {
			averageCS = cashSum / float32(countCS)
			averageCC = cardSum / float32(countCC)

			fmt.Printf("Average CS: %.2f\n", averageCS)
			fmt.Printf("Average CC: %.2f", averageCC)
			break
		}

	}

	if input == "End" {
		fmt.Println("Failed to collect required money for charity.")
	}
}
