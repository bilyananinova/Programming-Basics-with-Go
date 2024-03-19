package main

import (
	"fmt"
)

func main() {
	var a1, a2, n int
	fmt.Scanln(&a1)
	fmt.Scanln(&a2)
	fmt.Scanln(&n)

	for num1 := a1; num1 <= a2-1; num1++ {
		for num2 := 1; num2 <= n-1; num2++ {
			for num3 := 1; num3 <= n/2-1; num3++ {
				if num1%2 != 0 && (num2+num3+num1)%2 != 0 {
					fmt.Printf("%s-%d%d%d\n", string(num1), num2, num3, num1)
				}

			}
		}
	}
}
