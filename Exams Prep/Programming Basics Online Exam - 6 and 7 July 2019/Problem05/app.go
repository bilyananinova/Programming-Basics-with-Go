package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var sold, hearthstone, fornite, overwatch, others int
	var game string
	fmt.Scanln(&sold)

	scanner := bufio.NewScanner(os.Stdin)

	for s := 1; s <= sold; s++ {
		scanner.Scan()
		game = scanner.Text()

		if game == "Hearthstone" {
			hearthstone++
		} else if game == "Fornite" {
			fornite++
		} else if game == "Overwatch" {
			overwatch++
		} else {
			others++
		}
	}

	p1 := float64(hearthstone) / float64(sold) * 100
	p2 := float64(fornite) / float64(sold) * 100
	p3 := float64(overwatch) / float64(sold) * 100
	p4 := float64(others) / float64(sold) * 100

	fmt.Printf("Hearthstone - %.2f%%\n", p1)
	fmt.Printf("Fornite - %.2f%%\n", p2)
	fmt.Printf("Overwatch - %.2f%%\n", p3)
	fmt.Printf("Others - %.2f%%", p4)
}
