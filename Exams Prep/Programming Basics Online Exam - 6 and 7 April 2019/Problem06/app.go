package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var movie, ticketType string
	var seats, titalTickets, standardTickets, studentTickets, kidTickets int

	scanner.Scan()
	movie = scanner.Text()

	for movie != "Finish" {
		fmt.Scanln(&seats)
		fmt.Scanln(&ticketType)
		var sold int

		for ticketType != "End" {

			if ticketType == "standard" {
				standardTickets++
			} else if ticketType == "student" {
				studentTickets++
			} else if ticketType == "kid" {
				kidTickets++
			}

			sold++

			if seats == sold {
				break
			}

			fmt.Scanln(&ticketType)

			if ticketType == "Finish" {
				break
			}
		}

		titalTickets += sold
		fmt.Printf("%s - %.2f%% full.\n", movie, (float32(sold)/float32(seats))*100)

		if seats == sold {
			break
		}

		scanner.Scan()
		movie = scanner.Text()
	}

	fmt.Printf("Total tickets: %d\n", titalTickets)
	fmt.Printf("%.2f%% student tickets.\n", (float32(studentTickets)/float32(titalTickets))*100)
	fmt.Printf("%.2f%% standard tickets.\n", (float32(standardTickets)/float32(titalTickets))*100)
	fmt.Printf("%.2f%% kids tickets.", (float32(kidTickets) / float32(titalTickets) * 100))
}
