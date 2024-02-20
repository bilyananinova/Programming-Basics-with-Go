package main

import "fmt"

func main() {
	var hour int
	var day string

	fmt.Scanln(&hour)
	fmt.Scanln(&day)

	if hour >= 10 && hour <= 18 {
		if day == "Monday" || day == "Tuesday" || day == "Wednesday" || day == "Thursday" || day == "Friday" || day == "Saturday" {
			fmt.Println("open")
		} else if day == "Sunday" {
			fmt.Println("closed")
		}
	} else {
		fmt.Println("closed")
	}
}
