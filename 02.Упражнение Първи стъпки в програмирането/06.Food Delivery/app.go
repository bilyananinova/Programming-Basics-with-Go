package main

import "fmt"

func main() {
	var chickenMenu float32 = 10.35
	var fishMenu float32 = 12.40
	var vegeterianMenu float32 = 8.15
	var percentageDessert float32 = 20
	var deliveryPrice = 2.50

	var chickenMenusCount int
	var fishMenusCount int
	var vegeterianMenusCount int

	fmt.Scanln(&chickenMenusCount)
	fmt.Scanln(&fishMenusCount)
	fmt.Scanln(&vegeterianMenusCount)

	var chickenMenusPrice float32 = chickenMenu * float32(chickenMenusCount)
	var fishMenusPrice float32 = fishMenu * float32(fishMenusCount)
	var vegeterianMenusPrice float32 = vegeterianMenu * float32(vegeterianMenusCount)

	var menusPrice float32 = chickenMenusPrice + fishMenusPrice + vegeterianMenusPrice
	var dessertPrice float32 = menusPrice * (percentageDessert / 100)
	totalCosts := menusPrice + dessertPrice + float32(deliveryPrice)

	fmt.Println(totalCosts)
}
