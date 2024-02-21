package main

import "fmt"

func main() {
	var priceKgVegetables float32
	var priceKgFruits float32
	var kgsVegetables int
	var kgsFruits int

	fmt.Scanln(&priceKgVegetables)
	fmt.Scanln(&priceKgFruits)
	fmt.Scanln(&kgsVegetables)
	fmt.Scanln(&kgsFruits)

	total := (priceKgVegetables * float32(kgsVegetables)) + (priceKgFruits * float32(kgsFruits))
	total /= 1.94

	fmt.Printf("%.2f", total)
}
