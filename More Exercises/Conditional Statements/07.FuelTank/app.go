package main

import (
	"fmt"
	"strings"
)

func main() {
	var typeFuel string
	var qtyFuel float32

	fmt.Scan(&typeFuel, &qtyFuel)

	switch typeFuel {
	case "Diesel", "Gasoline", "Gas":
		if qtyFuel >= 25 {
			fmt.Printf("You have enough %s.", strings.ToLower(typeFuel))
		} else {
			fmt.Printf("Fill your tank with %s!", strings.ToLower(typeFuel))
		}
	default:
		fmt.Println("Invalid fuel!")
	}
}
