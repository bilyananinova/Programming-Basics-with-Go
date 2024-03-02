package main

import "fmt"

func main() {
	var h, m int

	for h <= 23 {

		for m <= 59 {
			fmt.Printf("%d : %d\n", h, m)
			m++
		}
		h++
		m = 0
	}
}
