package main

import "fmt"

func main() {
	var route string
	var tax float32
	var juniors, seniors, cyclists int

	fmt.Scan(&juniors, &seniors, &route)

	cyclists = juniors + seniors

	if route == "trail" {
		tax = float32(juniors)*5.50 + float32(seniors)*7

	} else if route == "cross-country" {
		tax = float32(juniors)*8.0 + float32(seniors)*9.50

		if cyclists >= 50 {
			tax *= 0.75
		}

	} else if route == "downhill" {
		tax = float32(juniors)*12.25 + float32(seniors)*13.75

	} else if route == "road" {
		tax = float32(juniors)*20 + float32(seniors)*21.50

	}

	tax *= 0.95

	fmt.Printf("%.2f", tax)
}

