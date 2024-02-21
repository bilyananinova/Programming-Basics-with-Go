package main

import "fmt"

func main() {
	var b1, b2, h float32

	fmt.Scanln(&b1)
	fmt.Scanln(&b2)
	fmt.Scanln(&h)

	area := (b1 + b2) * h / 2

	fmt.Printf("%.2f", area)
}
