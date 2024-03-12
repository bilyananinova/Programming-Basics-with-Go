package main

import "fmt"

func main() {
	var start, end, magicNumber, combinations int
	var isMagigCombination bool = false

	fmt.Scan(&start, &end, &magicNumber)

	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			combinations++

			if i+j == magicNumber {

				isMagigCombination = true

				fmt.Printf("Combination N:%d ", combinations)
				fmt.Printf("(%d + %d = %d)", i, j, magicNumber)
				break
			}
		}
		if isMagigCombination {
			break
		}
	}

	if !isMagigCombination {
		fmt.Printf("%d combinations - neither equals %d", combinations, magicNumber)
	}
}
