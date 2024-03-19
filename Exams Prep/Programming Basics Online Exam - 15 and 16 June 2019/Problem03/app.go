package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var movie, packet string
	var tickets int
	var price float32

	scanner.Scan()
	movie = scanner.Text()
	// fmt.Scanln(&movie)
	fmt.Scanln(&packet)
	fmt.Scanln(&tickets)

	if movie == "John Wick" {
		if packet == "Drink" {
			price = 12
		} else if packet == "Popcorn" {
			price = 15
		} else if packet == "Menu" {
			price = 19
		}
	} else if movie == "Star Wars" {
		if packet == "Drink" {
			price = 18
		} else if packet == "Popcorn" {
			price = 25
		} else if packet == "Menu" {
			price = 30
		}
	} else if movie == "Jumanji" {
		if packet == "Drink" {
			price = 9
		} else if packet == "Popcorn" {
			price = 11
		} else if packet == "Menu" {
			price = 14
		}
	}

	total := float32(tickets) * price

	if tickets >= 4 && movie == "Star Wars" {
		total *= 0.70
	}

	if tickets == 2 && movie == "Jumanji" {
		total *= 0.85
	}

	fmt.Printf("Your bill is %.2f leva.", total)
}
