package main

import (
	"fmt"
	"strconv"
)

func main() {
	var start, end int

	fmt.Scanln(&start)
	fmt.Scanln(&end)

	for i := start; i <= end; i++ {
		var number = strconv.Itoa(i)
		var oddSum, evenSum int

		for j := 0; j < len(number); j++ {
			num, _ := strconv.Atoi(string(number[j]))

			if j%2 != 0 {
				evenSum += num
			} else {
				oddSum += num
			}
		}

		if evenSum == oddSum {
			fmt.Printf("%s ", number)
		}
	}
}
