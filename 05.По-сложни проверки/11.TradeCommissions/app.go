package main

import "fmt"

func main() {
	var town string
	var sales, result, commission float32

	fmt.Scanln(&town)
	fmt.Scanln(&sales)

	switch town {
	case "Sofia":
		if sales > 10000 {
			commission = 0.12
		} else if sales > 1000 && sales <= 10000 {
			commission = 0.08
		} else if sales > 500 && sales <= 1000 {
			commission = 0.07
		} else if sales >= 0 && sales <= 500 {
			commission = 0.05
		} else {
			fmt.Println("error")
		}
	case "Varna":
		if sales > 10000 {
			commission = 0.13
		} else if sales > 1000 && sales <= 10000 {
			commission = 0.10
		} else if sales > 500 && sales <= 1000 {
			commission = 0.075
		} else if sales >= 0 && sales <= 500 {
			commission = 0.045
		} else {
			fmt.Println("error")
		}
	case "Plovdiv":
		if sales > 10000 {
			commission = 0.145
		} else if sales > 1000 && sales <= 10000 {
			commission = 0.12
		} else if sales > 500 && sales <= 1000 {
			commission = 0.08
		} else if sales >= 0 && sales <= 500 {
			commission = 0.055
		} else {
			fmt.Println("error")
		}
	default:
		fmt.Println("error")
	}

	if commission > 0 {
		result = sales * commission
		fmt.Printf("%.2f", result)
	}
}
