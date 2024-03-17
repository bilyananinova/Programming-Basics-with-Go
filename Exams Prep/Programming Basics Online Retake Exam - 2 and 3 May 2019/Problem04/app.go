package main

import "fmt"

func main() {
	var budget, price, costs float32
	var count, purchases int
	var input string

	fmt.Scanln(&budget)
	fmt.Scanln(&input)

	for input != "Stop" {
		fmt.Scanln(&price)
		count++

		if count == 3 {
			price /= 2
			count = 0
		}

		costs += price

		if budget < costs {
			fmt.Println("You don't have enough money!")
			fmt.Printf("You need %.2f leva!", costs-budget)
			break
		}

		purchases++

		fmt.Scanln(&input)
	}

	if input == "Stop" {
		fmt.Printf("You bought %d products for %.2f leva.", purchases, costs)
	}
}
