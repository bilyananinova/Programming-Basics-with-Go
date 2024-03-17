package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Scanln(&input)

	var primeSum, nonPrimeSum int

	for input != "stop" {
		number, _ := strconv.Atoi(input)
		var primeCounter int

		if number < 0 {
			fmt.Println("Number is negative.")
			fmt.Scanln(&input)
			continue
		}

		for i := 1; i <= number; i++ {
			if number%i == 0 {
				primeCounter++
			}
		}

		if primeCounter > 2 {
			nonPrimeSum += number
		} else {
			primeSum += number
		}

		fmt.Scanln(&input)
	}

	fmt.Printf("Sum of all prime numbers is: %d\n", primeSum)
	fmt.Printf("Sum of all non prime numbers is: %d", nonPrimeSum)
}
