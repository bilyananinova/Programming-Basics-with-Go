package main

import (
	"fmt"
	"strconv"
)

func main() {
	var poorGrades, count, index, sum, allProblems int
	var input, lastProblem string
	fmt.Scanln(&poorGrades)
	fmt.Scanln(&input)

	for input != "Enough" {

		if index%2 != 0 {
			grade, _ := strconv.Atoi(input)

			sum += grade

			if grade <= 4 {
				count++
			}

			if poorGrades == count {
				fmt.Printf("You need a break, %d poor grades.", count)
				break
			}

			allProblems++
		} else {
			lastProblem = input
		}

		index++

		fmt.Scanln(&input)

	}

	if input == "Enough" {
		fmt.Printf("Average score: %.2f\n", float32(sum)/float32(allProblems))
		fmt.Printf("Number of problems: %d\n", allProblems)
		fmt.Printf("Last problem: %s", lastProblem)
	}
}
