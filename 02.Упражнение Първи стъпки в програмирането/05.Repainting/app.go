package main

import "fmt"

func main() {
	var nylonPricePerSql float32 = 1.50
	var paintPricePerLitter float32 = 14.50
	var paintThinnerPricePerLittler float32 = 5.00
	var bags float32 = 0.40
	var percentageWork float32 = 30

	var nylonNeeded int
	var paintNeeded int
	var thinnerNeeded int
	var hoursWork int

	fmt.Scanln(&nylonNeeded)
	fmt.Scanln(&paintNeeded)
	fmt.Scanln(&thinnerNeeded)
	fmt.Scanln(&hoursWork)

	var costs float32 = (nylonPricePerSql * float32(nylonNeeded+2.0)) +
		(paintPricePerLitter * (float32(paintNeeded)*0.1 + float32(paintNeeded))) +
		(paintThinnerPricePerLittler * float32(thinnerNeeded)) + bags

	workersCost := (costs * percentageWork / 100) * float32(hoursWork)

	var totalCosts = costs + workersCost

	fmt.Println(totalCosts)
}
