package main

import (
	"fmt"
	"math"
)

func main() {
	var guest, budget int

	fmt.Scanln(&guest)
	fmt.Scanln(&budget)

	bitPrice := 4.0
	eggPrice := 0.45

	bits := math.Ceil(float64(guest) / 3)
	eggs := guest * 2

	total := bitPrice*bits + eggPrice*float64(eggs)

	diff := math.Abs(float64(budget) - total)

	if float64(budget) >= total {
		fmt.Printf("Lyubo bought %.0f Easter bread and %d eggs.\n", bits, eggs)
		fmt.Printf("He has %.2f lv. left.", diff)
	} else {
		fmt.Println("Lyubo doesn't have enough money.")
		fmt.Printf("He needs %.2f lv. more.", diff)
	}
}
