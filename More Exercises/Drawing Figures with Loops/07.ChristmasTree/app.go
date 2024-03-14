package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int

	fmt.Scanln(&n)

	for row := 0; row <= n; row++ {
		fmt.Println(strings.Repeat(" ", n+1) + "|" + strings.Repeat(" ", n+1))
		for col := 1; col <= n; col++ {
			fmt.Println(strings.Repeat(" ", n-col) + strings.Repeat("*", col) + " | " + strings.Repeat("*", col))
			continue
		}
		break
	}
}
