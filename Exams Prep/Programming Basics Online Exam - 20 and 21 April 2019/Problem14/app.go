package main

import "fmt"

func main() {
	var eggsInStore, eggs, soldEggs int
	var input string

	fmt.Scanln(&eggsInStore)
	fmt.Scanln(&input)

	for input != "Close" {
		fmt.Scanln(&eggs)

		if input == "Buy" {
			if eggsInStore < eggs {
				fmt.Println("Not enough eggs in store!")
				fmt.Printf("You can buy only %d.", eggsInStore)
				break
			}
			soldEggs += eggs
			eggsInStore -= eggs
		} else if input == "Fill" {
			eggsInStore += eggs
		}

		fmt.Scanln(&input)
	}

	if input == "Close" {
		fmt.Println("Store is closed!")
		fmt.Printf("%d eggs sold.", soldEggs)
	}
}
