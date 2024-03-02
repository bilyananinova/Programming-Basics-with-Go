package main

import "fmt"

func main() {
	var capacity, fens, aFens, bFens, vFens, gFens int
	var sector string

	fmt.Scanln(&capacity)
	fmt.Scanln(&fens)

	for i := 0; i < fens; i++ {
		fmt.Scanln(&sector)

		switch sector {
		case "A":
			aFens++
		case "B":
			bFens++
		case "V":
			vFens++
		case "G":
			gFens++
		}
	}

	aPercent := float32(aFens) / float32(fens) * 100
	bPercent := float32(bFens) / float32(fens) * 100
	vPercent := float32(vFens) / float32(fens) * 100
	gPercent := float32(gFens) / float32(fens) * 100
	allFensCapacity := float32(fens) / float32(capacity) * 100

	fmt.Printf("%.2f%%\n", aPercent)
	fmt.Printf("%.2f%%\n", bPercent)
	fmt.Printf("%.2f%%\n", vPercent)
	fmt.Printf("%.2f%%\n", gPercent)
	fmt.Printf("%.2f%%", allFensCapacity)
}
