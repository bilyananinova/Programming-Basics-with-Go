package main

import "fmt"

func main() {
	var result int

	fmt.Scanln(&result)

	if result >= 5 {
		fmt.Println("Excellent!")
	}
}
