package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			for k := 1; k <= 9; k++ {
				for l := 1; l <= 9; l++ {
					if n%i == 0 && n%j == 0 && n%k == 0 && n%l == 0 {
						fmt.Printf("%d%d%d%d ", i, j, k, l)
					}
				}
			}
		}
	}
}
