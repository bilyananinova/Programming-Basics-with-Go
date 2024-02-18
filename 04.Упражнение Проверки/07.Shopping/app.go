package main

import "fmt"

func main() {
	var budget float64
	var videoCardsCount, procesorsCount, ramCount int
	var videocardPrice float64 = 250

	fmt.Scanln(&budget)
	fmt.Scanln(&videoCardsCount)
	fmt.Scanln(&procesorsCount)
	fmt.Scanln(&ramCount)

	var procesorsPrice float64 = float64(videoCardsCount) * videocardPrice * 0.35
	var ramPrice float64 = float64(videoCardsCount) * videocardPrice * 0.10

	var sum float64 = float64(videoCardsCount)*videocardPrice + float64(procesorsCount)*procesorsPrice + float64(ramCount)*ramPrice

	if videoCardsCount > procesorsCount {
		sum = sum * 0.85
	}

	if sum <= budget {
		diff := budget - sum
		fmt.Printf("You have %.2f leva left!", diff)
	} else {
		diff := sum - budget
		fmt.Printf("Not enough money! You need %.2f leva more!", diff)
	}

}
