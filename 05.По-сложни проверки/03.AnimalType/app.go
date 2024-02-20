package main

import "fmt"

func main() {
	var animal string

	fmt.Scanln(&animal)

	switch animal {
	case "dog":
		fmt.Println("mammal")
	case "crocodile", "tortoise", "snake":
		fmt.Println("reptile")
	default:
		fmt.Println("unknown")
	}
}
