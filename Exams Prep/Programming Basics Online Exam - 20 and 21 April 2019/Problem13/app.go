package main

import "fmt"

func main() {
	var size, color string
	var count int
	var price float32

	fmt.Scanln(&size)
	fmt.Scanln(&color)
	fmt.Scanln(&count)

	if size == "Large" {
		if color == "Red" {
			price = 16
		} else if color == "Green" {
			price = 12
		} else if color == "Yellow" {
			price = 9
		}
	} else if size == "Medium" {
		if color == "Red" {
			price = 13
		} else if color == "Green" {
			price = 9
		} else if color == "Yellow" {
			price = 7
		}
	} else if size == "Small" {
		if color == "Red" {
			price = 9
		} else if color == "Green" {
			price = 8
		} else if color == "Yellow" {
			price = 5
		}
	}

	income := (price * float32(count)) * 0.65

	fmt.Printf("%.2f leva.", income)
}
