package main

import "fmt"

func main() {
	var days, hours int
	var sum, sumPerDay float32
	fmt.Scan(&days, &hours)

	for d := 1; d <= days; d++ {
		for h := 1; h <= hours; h++ {

			if d%2 == 0 && h%2 == 1 {
				sumPerDay += 2.50
			} else if d%2 == 1 && h%2 == 0 {
				sumPerDay += 1.25
			} else {
				sumPerDay += 1
			}

		}

		sum += sumPerDay
		fmt.Printf("Day: %d - %.2f leva\n", d, sumPerDay)
		sumPerDay = 0
	}

	fmt.Printf("Total: %.2f leva", sum)
}
