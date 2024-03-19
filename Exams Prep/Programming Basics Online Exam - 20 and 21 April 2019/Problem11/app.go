package main

import "fmt"

func main() {
	var flourPrice, flourKgs, sugarKgs float64
	var eggshells, yeastPackets int

	fmt.Scanln(&flourPrice)
	fmt.Scanln(&flourKgs)
	fmt.Scanln(&sugarKgs)
	fmt.Scanln(&eggshells)
	fmt.Scanln(&yeastPackets)

	sugarPrice := flourPrice * 0.75
	eggshellPrice := flourPrice * 1.10
	yeastPacketPrice := sugarPrice * 0.20

	total := flourKgs*flourPrice + sugarKgs*sugarPrice + float64(eggshells)*eggshellPrice + float64(yeastPackets)*yeastPacketPrice

	fmt.Printf("%.2f", total)
}
