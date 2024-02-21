package main

import (
	"fmt"
	"math"
)

func main() {
	var wInM float64
	var hInM float64

	fmt.Scanln(&wInM)
	fmt.Scanln(&hInM)

	hInCm := hInM*100 - 100
	desksRows := math.Floor(hInCm / 70)

	wInCm := wInM * 100
	rows := math.Floor(wInCm / 120)

	totalDesks := rows*desksRows - 3
	fmt.Println(totalDesks)

}
