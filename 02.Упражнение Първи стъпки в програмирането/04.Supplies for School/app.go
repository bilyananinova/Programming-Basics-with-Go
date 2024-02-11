package main

import "fmt"

func main() {
	var pricePackagePens float32 = 5.80
	var pricePackageMarkers float32 = 7.20
	var priceLiquid float32 = 1.20

	var packagePensCount int
	var packageMarkersCount int
	var liquidLitters int
	var discountPercent int

	fmt.Scanln(&packagePensCount)
	fmt.Scanln(&packageMarkersCount)
	fmt.Scanln(&liquidLitters)
	fmt.Scanln(&discountPercent)

	var totalPensPrice float32 = pricePackagePens * float32(packagePensCount)
	var totalMarkersPrice float32 = pricePackageMarkers * float32(packageMarkersCount)
	var totalLiquidPrice float32 = priceLiquid * float32(liquidLitters)

	var totalPrice = totalPensPrice + totalMarkersPrice + totalLiquidPrice

	var total float32 = totalPrice - (totalPrice * (float32(discountPercent) / 100))
	fmt.Println(total)
}
