package main

import (
	"fmt"
	"math"
)

func main() {
	var shape string
	var a float64
	var b float64
	var area float64

	fmt.Scanln(&shape)
	fmt.Scanln(&a)

	if shape == "rectangle" || shape == "triangle" {
		fmt.Scanln(&b)
	}

	if shape == "square" {
		area = a * a
	} else if shape == "rectangle" {
		area = a * b
	} else if shape == "circle" {
		area = math.Pi * math.Pow(a, 2)
	} else if shape == "triangle" {
		area = (a * b) / 2
	}

	fmt.Printf("%.3f", area)
}
