package main

import "fmt"

func main() {
	var startHour, startMinutes int
	var hours int

	fmt.Scanln(&startHour)
	fmt.Scanln(&startMinutes)

	minutes := startMinutes + 15

	if minutes >= 60 {
		hours = startHour + minutes/60
		minutes = minutes % 60

		if hours == 24 {
			hours = 0
		}

		if minutes < 10 {
			fmt.Printf("%d:0%d", hours, minutes)
		} else {
			fmt.Printf("%d:%d", hours, minutes)
		}

	} else {
		hours = startHour
		fmt.Printf("%d:%d", hours, minutes)
	}
}
