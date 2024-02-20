package main

import "fmt"

func main() {
	var number int

	fmt.Scanln(&number)

	if number != 0 && number < 100 || number > 200 {
		fmt.Println("invalid")
	}
}
