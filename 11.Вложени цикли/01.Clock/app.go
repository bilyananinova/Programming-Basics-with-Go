package main

import "fmt"

func main() {

	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {
			fmt.Printf("%d:%d\n", h, m)
		}
	}
}
