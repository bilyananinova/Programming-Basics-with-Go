package main

import "fmt"

func main() {
	var h, m, s int

	for h <= 23 {

		for m <= 59 {

			for s <= 59 {
				fmt.Printf("%d : %d : %d\n", h, m, s)
				s++
			}
			m++
			s = 0
		}
		h++
		m = 0
	}
}
