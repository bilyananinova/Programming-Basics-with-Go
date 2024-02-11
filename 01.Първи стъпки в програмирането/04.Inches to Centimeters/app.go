package main

import "fmt"

func main() {
	var inches int
	var centimetres float32 = 2.54

	fmt.Scanln(&inches)

	result := float32(inches) * centimetres

	fmt.Println(result)
}
