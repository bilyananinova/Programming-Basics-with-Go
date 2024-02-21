package main

import "fmt"

func main() {
	var celsius float32

	fmt.Scanln(&celsius)

	fahrenheit := celsius*1.8 + 32

	fmt.Printf("%.2f", fahrenheit)
}
