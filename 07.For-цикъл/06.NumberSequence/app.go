package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var currentNum, minNumber, maxNumber int32

	fmt.Scanln(&n)

	minNumber = math.MaxInt32
	maxNumber = math.MinInt32

	for i := 0; i < n; i++ {
		fmt.Scanln(&currentNum)

		if currentNum < minNumber {
			minNumber = currentNum
		}

		if currentNum > maxNumber {
			maxNumber = currentNum
		}
	}

	fmt.Printf("Max number: %d\n", maxNumber)
	fmt.Printf("Min number: %d", minNumber)
}
