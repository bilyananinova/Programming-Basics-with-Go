package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	var operator string

	fmt.Scan(&n1, &n2, &operator)

	if operator == "+" || operator == "-" || operator == "*" {
		var result int

		if operator == "+" {
			result = n1 + n2
		} else if operator == "-" {
			result = n1 - n2
		} else if operator == "*" {
			result = n1 * n2
		}

		if result%2 == 0 {
			fmt.Printf("%d %s %d = %d - even", n1, operator, n2, result)
		} else {
			fmt.Printf("%d %s %d = %d - odd", n1, operator, n2, result)
		}

	} else if operator == "/" || operator == "%" {

		if n2 == 0 {
			fmt.Printf("Cannot divide %d by zero", n1)
		} else {
			if operator == "/" {

				result := float32(n1) / float32(n2)
				fmt.Printf("%d %s %d = %.2f", n1, operator, n2, result)
			} else {
				result := n1 % n2
				fmt.Printf("%d %s %d = %d", n1, operator, n2, result)
			}
		}

	}
}
