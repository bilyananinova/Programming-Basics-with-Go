package main

import (
	"fmt"
	"math"
)

func main() {
	var magnoliasCount, hyacinthsCount, rosesCount, cactiCount int
	var giftPrice float32

	fmt.Scan(&magnoliasCount, &hyacinthsCount, &rosesCount, &cactiCount, &giftPrice)

	var magnolias float32 = 3.25
	var hyacinths float32 = 4
	var roses float32 = 3.50
	var cacti float32 = 8

	income := float32(magnoliasCount)*magnolias + float32(hyacinthsCount)*hyacinths + float32(rosesCount)*roses + float32(cactiCount)*cacti
	incomeWithoutTaxes := income * 0.95

	if incomeWithoutTaxes >= giftPrice {
		fmt.Printf("She is left with %.0f leva.", math.Floor(float64(incomeWithoutTaxes)-float64(giftPrice)))
	} else {
		fmt.Printf("She will have to borrow %.0f leva.", math.Ceil(float64(giftPrice)-float64(incomeWithoutTaxes)))
	}
}
