package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var result string
	fmt.Scan(&n)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			for k := 1; k <= 9; k++ {
				for l := 1; l <= 9; l++ {

					str := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k) + strconv.Itoa(l) + " "

					if i+j == k+l {
						if n%(i+j) == 0 {
							result += str
						}
					}

				}
			}
		}
	}

	fmt.Println(result)

}
