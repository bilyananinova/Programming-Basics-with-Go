package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	for row := 1; row <= n; row++ {

		fmt.Printf("%s ", "$")
		for col := 1; col < row; col++ {
			fmt.Printf("%s ", "$")
		}

		fmt.Println()
	}
}
