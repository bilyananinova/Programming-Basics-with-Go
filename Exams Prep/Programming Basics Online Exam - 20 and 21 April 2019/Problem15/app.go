package main

import (
	"fmt"
	"math"
)

func main() {

	var bits, usedSugar, usedFlour int
	var totalUsedSugar, totalUsedFlour float64
	var maxUsedSugar int32 = math.MinInt32
	var maxUsedFlour int32 = math.MinInt32

	fmt.Scanln(&bits)

	for b := 1; b <= bits; b++ {
		fmt.Scanln(&usedSugar)
		fmt.Scanln(&usedFlour)

		totalUsedSugar += float64(usedSugar)
		totalUsedFlour += float64(usedFlour)

		if maxUsedSugar < int32(usedSugar) {
			maxUsedSugar = int32(usedSugar)
		}

		if maxUsedFlour < int32(usedFlour) {
			maxUsedFlour = int32(usedFlour)
		}
	}

	var sugarPacket float64 = 950
	var flourPacket float64 = 750

	sugarPacks := math.Ceil(totalUsedSugar / sugarPacket)
	flourPacks := math.Ceil(totalUsedFlour / flourPacket)

	fmt.Printf("Sugar: %.0f\n", sugarPacks)
	fmt.Printf("Flour: %.0f\n", flourPacks)
	fmt.Printf("Max used flour is %d grams, max used sugar is %d grams.", maxUsedFlour, maxUsedSugar)
}
