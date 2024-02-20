package main

import "fmt"

func main() {
	var age float64
	var gender string

	fmt.Scanln(&age)
	fmt.Scanln(&gender)

	if gender == "m" {
		if age >= 16 {
			fmt.Println("Mr.")
		} else if age < 16 {
			fmt.Println("Master")
		}
	} else if gender == "f" {
		if age >= 16 {
			fmt.Println("Ms.")
		} else if age < 16 {
			fmt.Println("Miss")
		}
	}
}
