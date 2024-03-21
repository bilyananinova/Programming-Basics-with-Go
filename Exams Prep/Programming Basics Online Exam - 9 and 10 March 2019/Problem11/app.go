package main

import "fmt"

func main() {
	var taxForYear int

	fmt.Scanln(&taxForYear)

	shoesPrice := float64(taxForYear) * 0.60
	equipmentPrice := shoesPrice * 0.80
	ballPrice := equipmentPrice / 4
	accessoriesPrice := ballPrice / 5

	total := float64(taxForYear) + shoesPrice + equipmentPrice + ballPrice + accessoriesPrice

	fmt.Printf("%.2f", total)
}
