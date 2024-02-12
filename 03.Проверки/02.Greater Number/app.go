package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
