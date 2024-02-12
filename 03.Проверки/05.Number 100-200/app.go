package main

import "fmt"

func main() {
	var input int

	fmt.Scanln(&input)

	if input < 100 {
		fmt.Println("Less than 100")
	} else if input >= 100 && input <= 200 {
		fmt.Println("Between 100 and 200")
	} else {
		fmt.Println("Greater than 200")
	}
}
