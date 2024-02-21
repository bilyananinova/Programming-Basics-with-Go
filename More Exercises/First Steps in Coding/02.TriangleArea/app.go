package main

import "fmt"

func main() {
	var a, h float32

	fmt.Scanln(&a)
	fmt.Scanln(&h)

	area := a * h / 2

	fmt.Printf("%.2f", area)
}
