package main

import "fmt"

func main() {
	var mackerelPrice, sprinklePrice, bonitoKgs, safridKgs, musslesKgs float32

	fmt.Scan(&mackerelPrice, &sprinklePrice, &bonitoKgs, &safridKgs, &musslesKgs)

	bonitoPrice := mackerelPrice * 1.60
	safridPrice := sprinklePrice * 1.80

	totalBonito := bonitoKgs * bonitoPrice
	totalSafrid := safridKgs * safridPrice
	totalMussles := musslesKgs * 7.50

	total := totalBonito + totalSafrid + totalMussles

	fmt.Printf("%.2f", total)
}
