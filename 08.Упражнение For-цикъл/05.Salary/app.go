package main

import "fmt"

func main() {
	var tabs, salary int
	var tab string

	fmt.Scanln(&tabs)
	fmt.Scanln(&salary)

	for i := 0; i < tabs; i++ {

		fmt.Scanln(&tab)

		if tab == "Facebook" {
			salary -= 150
		} else if tab == "Instagram" {
			salary -= 100
		} else if tab == "Reddit" {
			salary -= 50
		}

	}

	if salary > 0 {
		fmt.Println(salary)
	}
	if salary <= 0 {
		fmt.Println("You have lost your salary.")
	}
}
