package main

import (
	"fmt"
	"math"
)

func main() {
	var startFirst, startSecond, diffFirst, diffSecond int
	fmt.Scan(&startFirst, &startSecond, &diffFirst, &diffSecond)

	endFirst := startFirst + diffFirst
	endSecond := startSecond + diffSecond

	for i := startFirst; i <= endFirst; i++ {
		for j := startSecond; j <= endSecond; j++ {

			isFirstNumPrime := true
			isSecondNumPrime := true

			for k := 2; k <= int(math.Sqrt(float64(i))); k++ {

				if i%k == 0 {
					isFirstNumPrime = false
					break
				}

			}

			for l := 2; l <= int(math.Sqrt(float64(j))); l++ {

				if j%l == 0 {
					isSecondNumPrime = false
					break
				}

			}

			if isFirstNumPrime && isSecondNumPrime {
				fmt.Printf("%d%d\n", i, j)
			}
		}
	}
}
