package main

import "fmt"

func main() {
	var degrees float32

	fmt.Scan(&degrees)

	if degrees >= 26 && degrees <= 35 {
		fmt.Println("Hot")
	} else if degrees >= 20.1 && degrees <= 25.9 {
		fmt.Println("Warm")
	} else if degrees >= 15 && degrees <= 20 {
		fmt.Println("Mild")
	} else if degrees >= 12 && degrees <= 14.9 {
		fmt.Println("Cool")
	} else if degrees >= 5 && degrees <= 11.9 {
		fmt.Println("Cold")
	} else {
		fmt.Println("unknown")
	}
}
