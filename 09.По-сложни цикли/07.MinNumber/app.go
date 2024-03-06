package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var input string
	var minNum = math.MaxInt32

	fmt.Scanln(&input)

	for input != "Stop" {
		num, _ := strconv.Atoi(input)

		if num < minNum {
			minNum = num
		}

		fmt.Scanln(&input)
	}

	fmt.Printf("%d", minNum)
}
