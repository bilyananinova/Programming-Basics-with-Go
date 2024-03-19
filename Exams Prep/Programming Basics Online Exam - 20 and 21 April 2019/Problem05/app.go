package main

import (
	"fmt"
	"math"
)

func main() {
	var eggsCount, r, o, b, g int
	var color string
	var maxEggs int32 = math.MinInt32
	var maxColor string

	fmt.Scanln(&eggsCount)

	for e := 1; e <= eggsCount; e++ {
		fmt.Scanln(&color)

		if color == "red" {
			r++
		} else if color == "orange" {
			o++
		} else if color == "blue" {
			b++
		} else if color == "green" {
			g++
		}

		if maxEggs < int32(r) && color == "red" {
			maxEggs = int32(r)
			maxColor = color
		} else if maxEggs < int32(o) && color == "orange" {
			maxEggs = int32(o)
			maxColor = color
		} else if maxEggs < int32(b) && color == "blue" {
			maxEggs = int32(b)
			maxColor = color
		} else if maxEggs < int32(g) && color == "green" {
			maxEggs = int32(g)
			maxColor = color
		}
	}

	fmt.Printf("Red eggs: %d\n", r)
	fmt.Printf("Orange eggs: %d\n", o)
	fmt.Printf("Blue eggs: %d\n", b)
	fmt.Printf("Green eggs: %d\n", g)
	fmt.Printf("Max eggs: %d -> %s", maxEggs, maxColor)
}
