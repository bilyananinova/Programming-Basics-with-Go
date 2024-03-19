package main

import (
	"fmt"
	"math"
)

func main() {
	var guests int
	var platePrice, budget float64

	fmt.Scanln(&guests)
	fmt.Scanln(&platePrice)
	fmt.Scanln(&budget)

	if guests >= 10 && guests <= 15 {
		platePrice *= 0.85
	} else if guests > 15 && guests <= 20 {
		platePrice *= 0.80
	} else if guests > 20 {
		platePrice *= 0.75
	}

	cakePrice := budget * 0.10

	total := cakePrice + (platePrice * float64(guests))

	diff := math.Abs(budget - total)

	if budget >= total {
		fmt.Printf("It is party time! %.2f leva left.", diff)
	} else {
		fmt.Printf("No party! %.2f leva needed.", diff)
	}

}
