package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scanln(&n)

	for i := 0; i <= n; i += 2 {
		i := math.Pow(2, float64(i))
		fmt.Println(i)
	}
}
