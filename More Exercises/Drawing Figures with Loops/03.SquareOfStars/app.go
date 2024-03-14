package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s ", "*")
		}
		fmt.Println()
	}
}
