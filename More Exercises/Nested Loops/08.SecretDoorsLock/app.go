package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	for i := 1; i <= a; i++ {
		for k := 1; k <= b; k++ {
			for l := 1; l <= c; l++ {

				if i%2 == 0 && l%2 == 0 {
					if k == 2 || k == 3 || k == 5 || k == 7 {
						str := strconv.Itoa(i) + " " + strconv.Itoa(k) + " " + strconv.Itoa(l)
						fmt.Println(str)
					}
				}
			}
		}
	}
}
