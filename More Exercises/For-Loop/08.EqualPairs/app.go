package main

import (
	"fmt"
	"math"
)

func main() {
	var pairs, numA, numB, currentSum, sum int
	var maxDiff int32 = math.MinInt32

	fmt.Scanln(&pairs)

	for i := 0; i < pairs*2; i += 2 {

		fmt.Scanln(&numA)
		fmt.Scanln(&numB)

		currentSum = numA + numB

		if sum != currentSum {
			currentDiff := math.Abs(float64(sum) - float64(currentSum))

			if int32(currentDiff) > maxDiff && i != 0 {
				maxDiff = int32(currentDiff)
			}

			sum = currentSum
		}

	}

	if maxDiff > 0 {
		fmt.Printf("No, maxdiff=%d", maxDiff)
	} else {
		fmt.Printf("Yes, value=%d", sum)
	}

}
