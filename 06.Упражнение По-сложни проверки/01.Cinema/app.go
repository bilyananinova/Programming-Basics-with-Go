package main

import "fmt"

func main() {
	var movie string
	var r int
	var c int
	var ticketPrice float32

	fmt.Scanln(&movie)
	fmt.Scanln(&r)
	fmt.Scanln(&c)

	if movie == "Premiere" {
		ticketPrice = 12.00
	} else if movie == "Normal" {
		ticketPrice = 7.50
	} else if movie == "Discount" {
		ticketPrice = 5.0
	}

	total := float32(r) * float32(c) * ticketPrice

	fmt.Printf("%.2f leva", total)
}
