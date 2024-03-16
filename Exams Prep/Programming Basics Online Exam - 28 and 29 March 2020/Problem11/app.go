package main

import "fmt"

func main() {
	var bitcoints int
	var uan, commision float64

	fmt.Scanln(&bitcoints)
	fmt.Scanln(&uan)
	fmt.Scanln(&commision)

	bitcointsInLv := bitcoints * 1168
	uansInUSD := uan * 0.15
	usdInLv := uansInUSD * 1.76

	moneyInLv := float64(bitcointsInLv) + usdInLv
	moneyInEuro := moneyInLv / 1.95
	result := moneyInEuro - (moneyInEuro * commision / 100)

	fmt.Printf("%.2f", result)
}
