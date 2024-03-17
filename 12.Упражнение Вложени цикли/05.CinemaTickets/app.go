package main

import "fmt"

func main() {
	var input, typeOfticket string
	var capacity, standardCount, studentCount, kidCount, totalTickets int

	fmt.Scanln(&input)

	for input != "Finish" {
		fmt.Scanln(&capacity)
		fmt.Scanln(&typeOfticket)
		var soldTicketsPerCurrentMovie int = 0

		for typeOfticket != "End" {

			if typeOfticket == "standard" {
				standardCount++
			} else if typeOfticket == "student" {
				studentCount++
			} else if typeOfticket == "kid" {
				kidCount++
			}
			soldTicketsPerCurrentMovie++
			totalTickets++

			if capacity == soldTicketsPerCurrentMovie {
				break
			}

			fmt.Scanln(&typeOfticket)
		}

		fmt.Printf("%s - %.2f%% full.\n", input, float32(soldTicketsPerCurrentMovie)/float32(capacity)*100)
		fmt.Scanln(&input)
	}
	fmt.Printf("Total tickets: %d\n", totalTickets)
	fmt.Printf("%.2f%% student tickets.\n", float32(studentCount)/float32(totalTickets)*100)
	fmt.Printf("%.2f%% standard tickets.\n", float32(standardCount)/float32(totalTickets)*100)
	fmt.Printf("%.2f%% kids tickets.", float32(kidCount)/float32(totalTickets)*100)
}
