package main

import "fmt"

func main() {
	var input string
	var budget, money, sum float32
	fmt.Scanln(&input)

	for input != "End" {
		fmt.Scanln(&budget)
		for sum < budget {
			fmt.Scanln(&money)
			sum += money

			if sum >= budget {
				fmt.Printf("Going to %s!\n", input)
			}
		}
		fmt.Scanln(&input)
		sum = 0
	}
}
