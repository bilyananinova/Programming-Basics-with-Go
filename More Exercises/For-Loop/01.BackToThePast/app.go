package main

import (
	"fmt"
	"math"
)

func main() {
	var inheritanceMoney float64
	var endYear int
	var ivansYears int = 18

	fmt.Scan(&inheritanceMoney, &endYear)

	for i := 1800; i <= endYear; i++ {

		inheritanceMoney -= 12000.00

		if i%2 != 0 {
			inheritanceMoney -= (float64(ivansYears) * 50)
		}

		ivansYears++
	}

	if inheritanceMoney >= 0 {
		fmt.Printf("Yes! He will live a carefree life and will have %.2f dollars left.", inheritanceMoney)
	} else {
		fmt.Printf("He will need %.2f dollars to survive.", math.Abs(inheritanceMoney))
	}

}
