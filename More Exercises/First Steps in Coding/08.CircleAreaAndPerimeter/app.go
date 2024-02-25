package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Scan(&r)

	area := math.Pi * math.Pow(r, 2)
	result := 2 * (math.Pi * r)

	fmt.Printf("%.2f\n", area)
	fmt.Printf("%.2f", result)
}
