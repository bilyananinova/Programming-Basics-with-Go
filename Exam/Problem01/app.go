package main

import "fmt"

func main() {
	var group, nights, cards, tickets int
	fmt.Scanln(&group)
	fmt.Scanln(&nights)
	fmt.Scanln(&cards)
	fmt.Scanln(&tickets)

	nightPrice := 20.0
	cardPrice := 1.60
	ticketPrice := 6.0

	nightsAmount := nightPrice * float64(nights)
	cardsAmount := cardPrice * float64(cards)
	ticketsAmount := ticketPrice * float64(tickets)
	totalPerson := nightsAmount + cardsAmount + ticketsAmount
	totalGroup := totalPerson * float64(group)
	additionalCosts := totalGroup * 0.25

	totalCosts := totalGroup + additionalCosts
	fmt.Printf("%.2f", totalCosts)

}
