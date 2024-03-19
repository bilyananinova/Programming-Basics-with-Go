package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var bitsCount int
	var baker, input, maxBaker string
	var maxPoints int32 = math.MinInt32

	fmt.Scanln(&bitsCount)

	for b := 1; b <= bitsCount; b++ {
		scanner.Scan()
		baker = scanner.Text()
		// fmt.Scanln(&baker)
		fmt.Scanln(&input)
		var points int

		for input != "Stop" {
			p, _ := strconv.Atoi(input)

			points += p

			fmt.Scanln(&input)
		}

		fmt.Printf("%s has %d points.\n", baker, points)

		if maxPoints < int32(points) {
			maxPoints = int32(points)
			maxBaker = baker
			fmt.Printf("%s is the new number 1!\n", baker)
		}
	}

	fmt.Printf("%s won competition with %d points!", maxBaker, maxPoints)
}
