package main

import "fmt"

func main() {
	var days, hours int
	var totalTax, taxPerDay float32

	fmt.Scanln(&days)
	fmt.Scanln(&hours)

	for d := 1; d <= days; d++ {

		for h := 1; h <= hours; h++ {

			if d%2 == 0 && h%2 == 1 {
				taxPerDay += 2.50
			} else if d%2 == 1 && h%2 == 0 {
				taxPerDay += 1.25
			} else {
				taxPerDay += 1
			}
		}

		fmt.Printf("Day: %d - %.2f leva\n", d, taxPerDay)

		totalTax += taxPerDay
		taxPerDay = 0
	}

	fmt.Printf("Total: %.2f leva", totalTax)
}
