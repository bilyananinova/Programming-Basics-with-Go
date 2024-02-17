package main

import (
	"fmt"
	"math"
)

func main() {

	var tripPrice float64
	var puzzleCount, dollCount, bearCount, minionsCount, truckCount int

	fmt.Scanln(&tripPrice)
	fmt.Scanln(&puzzleCount)
	fmt.Scanln(&dollCount)
	fmt.Scanln(&bearCount)
	fmt.Scanln(&minionsCount)
	fmt.Scanln(&truckCount)

	puzzlePrice := 2.60
	dollPrice := 3.0
	bearPrice := 4.10
	minionsPrice := 8.20
	truckPrice := 2.0

	sum := (float64(puzzleCount) * puzzlePrice) + (float64(dollCount) * dollPrice) + (float64(bearCount) * bearPrice) + (float64(minionsCount) * minionsPrice) + (float64(truckCount) * truckPrice)
	toys := puzzleCount + dollCount + bearCount + minionsCount + truckCount

	if toys >= 50 {
		sum = sum * 0.75
	}

	profit := sum * 0.90
	diff := math.Abs(profit - tripPrice)
	if profit >= tripPrice {
		fmt.Printf("Yes! %.2f lv left.", diff)
	} else {
		fmt.Printf("Not enough money! %.2f lv needed.", diff)
	}
}
