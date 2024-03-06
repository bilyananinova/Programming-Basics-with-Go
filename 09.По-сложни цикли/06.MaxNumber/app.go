package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var input string
	var maxNum = math.MinInt32

	fmt.Scanln(&input)

	for input != "Stop" {
		num, _ := strconv.Atoi(input)

		if num > maxNum {
			maxNum = num
		}

		fmt.Scanln(&input)
	}

	fmt.Printf("%d", maxNum)
}
