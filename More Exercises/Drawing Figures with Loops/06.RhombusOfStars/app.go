package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int

	fmt.Scanln(&n)

	for row := 1; row <= n; row++ {
		space := strings.Repeat(" ", n-row)
		fmt.Printf("%s", space)
		fmt.Print("*")
		str := strings.Repeat(" *", row-1)
		fmt.Println(str)
	}

	for i := n - 1; i >= 0; i-- {
		space := strings.Repeat(" ", n-i)
		fmt.Printf("%s", space)
		str := strings.Repeat("* ", i)
		fmt.Println(str)
	}

}
