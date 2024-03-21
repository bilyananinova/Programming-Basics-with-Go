package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var movie, typeHall string
	var tickets int
	var price float64

	scanner.Scan()
	movie = scanner.Text()
	scanner.Scan()
	typeHall = scanner.Text()
	fmt.Scanln(&tickets)

	if movie == "A Star Is Born" {
		if typeHall == "normal" {
			price = 7.50
		} else if typeHall == "luxury" {
			price = 10.50
		} else if typeHall == "ultra luxury" {
			price = 13.50
		}
	} else if movie == "Bohemian Rhapsody" {
		if typeHall == "normal" {
			price = 7.35
		} else if typeHall == "luxury" {
			price = 9.45
		} else if typeHall == "ultra luxury" {
			price = 12.75
		}
	} else if movie == "Green Book" {
		if typeHall == "normal" {
			price = 8.15
		} else if typeHall == "luxury" {
			price = 10.25
		} else if typeHall == "ultra luxury" {
			price = 13.25
		}
	} else if movie == "The Favourite" {
		if typeHall == "normal" {
			price = 8.75
		} else if typeHall == "luxury" {
			price = 11.55
		} else if typeHall == "ultra luxury" {
			price = 13.95
		}
	}

	total := price * float64(tickets)

	fmt.Printf("%s -> %.2f lv.", movie, total)
}
