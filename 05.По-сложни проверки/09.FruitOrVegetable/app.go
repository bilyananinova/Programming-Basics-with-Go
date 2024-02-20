package main

import "fmt"

func main() {
	var product string

	fmt.Scanln(&product)

	switch product {
	case "banana", "apple", "kiwi", "cherry", "lemon", "grapes":
		fmt.Println("fruit")
	case "tomato", "cucumber", "pepper", "carrot":
		fmt.Println("vegetable")
	default:
		fmt.Println("unknown")
	}
}
