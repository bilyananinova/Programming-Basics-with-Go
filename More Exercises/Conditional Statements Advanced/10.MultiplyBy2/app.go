package main

import "fmt"

func main() {
	var n, result float32

	fmt.Scan(&n)

	for n >= 0 {
		result = n * 2
		if result >= 0 {
			fmt.Printf("Result: %.2f\n", result)
		}

		fmt.Scan(&n)
	}

	fmt.Println("Negative number!")

}
