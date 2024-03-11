package main

import "fmt"

func main() {
	var start, end, magicNum, combinationNum, sum int
	fmt.Scan(&start, &end, &magicNum)

	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {

			sum = i + j
			combinationNum++

			if sum == magicNum {
				fmt.Printf("Combination N:%d (%d + %d = %d)", combinationNum, i, j, i+j)
				break
			}
		}

		if sum == magicNum {
			break
		}
	}

	if sum != magicNum {
		fmt.Printf("%d combinations - neither equals %d", combinationNum, magicNum)
	}
}
