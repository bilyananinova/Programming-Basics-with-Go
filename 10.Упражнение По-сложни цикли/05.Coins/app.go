package main

import (
	"fmt"
	"math"
)

func main() {
	var change, changeStotinki float64
	var count int

	fmt.Scanln(&change)

	changeStotinki = math.Floor(change * 100)

	for changeStotinki > 0 {
		count++
		if changeStotinki >= 200 {
			changeStotinki -= 200
		} else if changeStotinki >= 100 {
			changeStotinki -= 100
		} else if changeStotinki >= 50 {
			changeStotinki -= 50
		} else if changeStotinki >= 20 {
			changeStotinki -= 20
		} else if changeStotinki >= 10 {
			changeStotinki -= 10
		} else if changeStotinki >= 5 {
			changeStotinki -= 5
		} else if changeStotinki >= 2 {
			changeStotinki -= 2
		} else if changeStotinki >= 1 {
			changeStotinki -= 1
		}
	}

	fmt.Println(count)
}
