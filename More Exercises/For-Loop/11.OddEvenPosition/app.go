package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var num, oddSum, evenSum float32
	var minOdd float32 = math.MaxInt32
	var maxOdd float32 = math.MinInt32
	var minEven float32 = math.MaxInt32
	var maxEven float32 = math.MinInt32

	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&num)

		if i%2 != 0 {
			oddSum += num

			if maxOdd < num {
				maxOdd = num
			}

			if num < minOdd {
				minOdd = num
			}
		} else {
			evenSum += num

			if maxEven < num {
				maxEven = num
			}

			if num < minEven {
				minEven = num
			}
		}
	}

	fmt.Printf("OddSum=%.2f,\n", oddSum)

	if minOdd != math.MaxInt32 {
		fmt.Printf("OddMin=%.2f,\n", minOdd)
	} else {
		fmt.Println("OddMin=No,")
	}

	if maxOdd != math.MinInt32 {
		fmt.Printf("OddMax=%.2f,\n", maxOdd)
	} else {
		fmt.Println("OddMax=No,")
	}

	fmt.Printf("EvenSum=%.2f,\n", evenSum)

	if minEven != math.MaxInt32 {
		fmt.Printf("EvenMin=%.2f,\n", minEven)
	} else {
		fmt.Println("EvenMin=No,")
	}

	if maxEven != math.MinInt32 {
		fmt.Printf("EvenMax=%.2f\n", maxEven)
	} else {
		fmt.Println("EvenMax=No")
	}
}
