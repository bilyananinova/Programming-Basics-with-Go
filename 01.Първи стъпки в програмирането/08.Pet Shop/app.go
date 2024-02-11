package main

import "fmt"

func main() {
	var dogsFood float32 = 2.50
	var catsFood float32 = 4.0
	var dogsFoodCount int
	var catsFoodCount int

	fmt.Scanln(&dogsFoodCount)
	fmt.Scanln(&catsFoodCount)

	costs := dogsFood*float32(dogsFoodCount) + catsFood*float32(catsFoodCount)

	fmt.Printf("%f lv.", costs)

}
