package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	var width, height int
	fmt.Scanln(&width)
	fmt.Scanln(&height)

	cake := width * height

	for input != "STOP" {
		fmt.Scanln(&input)
		pices, _ := strconv.Atoi(input)

		cake -= pices

		if cake <= 0 {
			fmt.Printf("No more cake left! You need %.0f pieces more.", math.Abs(float64(cake)))
			break
		}
	}

	if input == "STOP" {
		fmt.Printf("%d pieces are left.", cake)
	}
}
