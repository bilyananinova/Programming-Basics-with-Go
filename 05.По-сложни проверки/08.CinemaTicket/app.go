package main

import "fmt"

func main() {
	var day string

	fmt.Scanln(&day)

	switch day {
	case "Monday", "Tuesday", "Friday":
		fmt.Println(12)
	case "Wednesday", "Thursday":
		fmt.Println(14)
	case "Saturday", "Sunday":
		fmt.Println(16)
	}
}
