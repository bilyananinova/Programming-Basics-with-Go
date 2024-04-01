package main

import "fmt"

func main() {
	var tishirtPrice, limitAmount float32

	fmt.Scanln(&tishirtPrice)
	fmt.Scanln(&limitAmount)

	shortsPrice := tishirtPrice * 0.75
	socksPrice := shortsPrice * 0.20
	shoesPrice := (tishirtPrice + shortsPrice) * 2
	totalSum := tishirtPrice + shortsPrice + socksPrice + shoesPrice
	totalAfterDiscount := totalSum * 0.85

	if totalAfterDiscount >= limitAmount {
		fmt.Println("Yes, he will earn the world-cup replica ball!")
		fmt.Printf("His sum is %.2f lv.", totalAfterDiscount)
	} else {
		neededSum := limitAmount - totalAfterDiscount
		fmt.Println("No, he will not earn the world-cup replica ball.")
		fmt.Printf("He needs %.2f lv. more.", neededSum)
	}
}
