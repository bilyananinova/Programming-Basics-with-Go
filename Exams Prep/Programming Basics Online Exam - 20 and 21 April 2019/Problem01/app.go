package main

import "fmt"

func main() {
	var bits, eggshells, cookies int

	fmt.Scanln(&bits)
	fmt.Scanln(&eggshells)
	fmt.Scanln(&cookies)

	bitPrice := 3.20
	eggshellPrice := 4.35
	cookiesPrice := 5.40
	eggPaintPrice := 0.15

	totalBitsPrice := float64(bits) * bitPrice
	totalEggsShellPrice := (float64(eggshells) * eggshellPrice) + (float64(eggshells) * 12 * eggPaintPrice)
	totalCookiesPrice := float64(cookies) * cookiesPrice

	result := totalBitsPrice + totalEggsShellPrice + totalCookiesPrice

	fmt.Printf("%.2f", result)
}
