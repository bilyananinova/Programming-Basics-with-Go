package main

import "fmt"

func main() {
	var shoesPercentage float32 = 40
	var equipmentPercentage float32 = 20
	var taxPerYear float32

	fmt.Scanln(&taxPerYear)

	var shoesPrice float32 = taxPerYear - (taxPerYear * shoesPercentage / 100)
	var equipmentPrice float32 = shoesPrice - (shoesPrice * equipmentPercentage / 100)
	var ballPrice float32 = equipmentPrice / 4
	var accessoriesPrice float32 = ballPrice / 5

	var totalCosts = taxPerYear + shoesPrice + equipmentPrice + ballPrice + accessoriesPrice
	fmt.Println(totalCosts)
}
