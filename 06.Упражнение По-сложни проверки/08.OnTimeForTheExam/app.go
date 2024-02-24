package main

import "fmt"

func main() {
	var examHour, examMinutes, arrivalHour, arrivalMinutes int

	fmt.Scan(&examHour, &examMinutes, &arrivalHour, &arrivalMinutes)

	examInMinutes := examHour*60 + examMinutes
	arrivalInMinutes := arrivalHour*60 + arrivalMinutes

	if arrivalInMinutes > examInMinutes {
		fmt.Println("Late")
		diff := arrivalInMinutes - examInMinutes

		if diff >= 60 {
			hour := diff / 60
			mins := diff % 60

			fmt.Printf("%d:%.2d hours after the start", hour, mins)
		} else {
			fmt.Printf("%d minutes after the start", diff)
		}
	} else if arrivalInMinutes <= examInMinutes {
		diff := examInMinutes - arrivalInMinutes

		if arrivalInMinutes == examInMinutes || diff <= 30 {
			fmt.Println("On time")

			if diff != 0 {
				fmt.Printf("%d minutes before the start", diff)
			}

		} else if arrivalInMinutes < examInMinutes {
			fmt.Println("Early")

			if diff >= 60 {
				hour := diff / 60
				mins := diff % 60

				fmt.Printf("%d:%.2d hours before the start", hour, mins)
			} else {
				fmt.Printf("%d minutes before the start", diff)
			}

		}
	}

}
