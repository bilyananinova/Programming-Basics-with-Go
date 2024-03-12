package main

import (
	"fmt"
	"strconv"
)

func main() {
	var start, end int
	var result string
	fmt.Scan(&start)
	fmt.Scan(&end)

	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			for k := start; k <= end; k++ {
				for l := start; l <= end; l++ {
					str := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k) + strconv.Itoa(l) + " "
					if (i%2 == 0 && l%2 == 1) || (i%2 == 1 && l%2 == 0) {
						if i > l {
							if (j+k)%2 == 0 {
								result += str
							}
						}
					}

				}
			}
		}
	}

	fmt.Println(result)
}
