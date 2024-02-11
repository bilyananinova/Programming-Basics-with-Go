package main

import "fmt"

func main() {
	var deposite float32
	var term float32
	var interest float32

	fmt.Scanln(&deposite)
	fmt.Scanln(&term)
	fmt.Scanln(&interest)

	var interestPerMonth = (deposite * interest / 100) / 12
	interestPerTerm := term * interestPerMonth
	result := deposite + interestPerTerm

	fmt.Println(result)
}
