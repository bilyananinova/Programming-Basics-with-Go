package main

import "fmt"

func main() {
	var sqmPrice float32 = 7.61
	var discountPercentage = 18

	var yard float32

	fmt.Scanln(&yard)

	var price float32 = yard * sqmPrice
	var discount float32 = float32(discountPercentage) / 100 * float32(price)
	var finalPrice float32 = price - discount

	fmt.Printf("The final price is: %f lv.\n", finalPrice)
	fmt.Printf("The discount is: %f lv.", discount)
}
