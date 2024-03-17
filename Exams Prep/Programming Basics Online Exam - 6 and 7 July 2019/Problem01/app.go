package main

import (
	"fmt"
	"math"
)

func main() {
	var people int
	var tax, chairPrice, umbrelaPrice float64

	fmt.Scanln(&people)
	fmt.Scanln(&tax)
	fmt.Scanln(&chairPrice)
	fmt.Scanln(&umbrelaPrice)

	allTaxes := float64(people) * tax
	umbrelas := math.Ceil(float64(people) / 2)
	umbrelasTaxes := umbrelas * umbrelaPrice
	chairTaxes := math.Ceil(float64(people)*0.75) * chairPrice

	fmt.Printf("%.2f lv.", allTaxes+umbrelasTaxes+chairTaxes)
}
