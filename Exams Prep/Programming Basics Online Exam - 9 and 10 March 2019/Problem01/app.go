package main

import (
	"fmt"
	"math"
)

func main() {
	var racketsPrice float64
	var rackets, shoes int

	fmt.Scanln(&racketsPrice)
	fmt.Scanln(&rackets)
	fmt.Scanln(&shoes)

	totalRackets := racketsPrice * float64(rackets)
	totalShoes := (racketsPrice / 6) * float64(shoes)
	total := totalRackets + totalShoes
	otherEquipment := total * 0.20
	total += otherEquipment

	paidByPlayer := math.Floor(total / 8)
	fmt.Printf("Price to be paid by Djokovic %.0f\n", paidByPlayer)

	paidBySponsors := math.Ceil(total / 8 * 7)
	fmt.Printf("Price to be paid by sponsors %.0f\n", paidBySponsors)

}
