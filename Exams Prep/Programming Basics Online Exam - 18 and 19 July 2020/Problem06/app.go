package main

import (
	"fmt"
)

func main() {
	var start, end int

	fmt.Scanln(&start)
	fmt.Scanln(&end)

	thousandsStartNumber := start / 1000
	hundredsStartNumber := start / 100 % 10
	tensStartNumber := start / 10 % 10
	unitsStartNumber := start / 1 % 10

	thousandsEndNumber := end / 1000
	hundredsEndNumber := end / 100 % 10
	tensEndNumber := end / 10 % 10
	unitsEndNumber := end / 1 % 10

	for thousands := thousandsStartNumber; thousands <= thousandsEndNumber; thousands++ {
		for hundreds := hundredsStartNumber; hundreds <= hundredsEndNumber; hundreds++ {
			for tens := tensStartNumber; tens <= tensEndNumber; tens++ {
				for units := unitsStartNumber; units <= unitsEndNumber; units++ {

					if thousands%2 != 0 && hundreds%2 != 0 && tens%2 != 0 && units%2 != 0 {
						fmt.Printf("%d%d%d%d ", thousands, hundreds, tens, units)
					}
				}
			}
		}
	}
}
