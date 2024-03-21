package main

import (
	"fmt"
)

func main() {
	var target, highJump, jumps int
	var isFailed bool = false
	fmt.Scanln(&target)

	start := target - 30

	for start <= target {
		for i := 0; i < 3; i++ {
			fmt.Scanln(&highJump)

			jumps++

			if highJump > start {
				start += 5
				break
			}

			if i == 2 {
				fmt.Printf("Tihomir failed at %dcm after %d jumps.", highJump, jumps)
				isFailed = true
				break
			}

		}

		if isFailed {
			break
		}

	}

	if !isFailed {
		fmt.Printf("Tihomir succeeded, he jumped over %dcm after %d jumps.", target, jumps)
	}
}
