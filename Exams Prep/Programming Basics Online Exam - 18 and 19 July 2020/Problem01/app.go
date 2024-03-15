package main

import "fmt"

func main() {
	var name string
	var ticketsAdult, ticketsKids int
	var ticketPriceAdult, service, profit float64

	fmt.Scanln(&name)
	fmt.Scanln(&ticketsAdult)
	fmt.Scanln(&ticketsKids)
	fmt.Scanln(&ticketPriceAdult)
	fmt.Scanln(&service)

	ticketPriceKids := float64(ticketPriceAdult) - (float64(ticketPriceAdult) * 0.70)
	ticketPriceAdultService := ticketPriceAdult + service
	ticketPriceKidsService := ticketPriceKids + service

	profit = ((float64(ticketsAdult) * ticketPriceAdultService) + (float64(ticketsKids) * ticketPriceKidsService)) * 0.20

	fmt.Printf("The profit of your agency from %s tickets is %.2f lv.", name, profit)
}
