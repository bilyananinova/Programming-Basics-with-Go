package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var film string
	var days, tickets, cinemaPercent int
	var ticketsPrice float64

	scanner.Scan()
	film = scanner.Text()

	// fmt.Scanln(&film)
	fmt.Scanln(&days)
	fmt.Scanln(&tickets)
	fmt.Scanln(&ticketsPrice)
	fmt.Scanln(&cinemaPercent)

	totalTickets := float64(tickets) * ticketsPrice * float64(days)
	totalIncome := totalTickets - (totalTickets * float64(cinemaPercent) / 100)

	fmt.Printf("The profit from the movie %s is %.2f lv.", film, totalIncome)

}
