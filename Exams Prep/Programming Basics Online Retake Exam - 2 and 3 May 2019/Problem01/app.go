package main

import "fmt"

func main() {
	var strawberriesPrice, bananasKg, orangesKg, raspberriesKg, strawberriesKg float32
	fmt.Scan(&strawberriesPrice, &bananasKg, &orangesKg, &raspberriesKg, &strawberriesKg)

	raspberriesPrice := strawberriesPrice / 2
	orangesPrice := raspberriesPrice * 0.60
	bananasPrice := raspberriesPrice * 0.20

	total := strawberriesKg*strawberriesPrice + bananasKg*bananasPrice + orangesKg*orangesPrice + raspberriesKg*raspberriesPrice

	fmt.Printf("%.2f", total)
}
