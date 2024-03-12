package main

import (
	"fmt"
	"strconv"
)

func main() {
	var result string
	var x, y, max, count int
	var symA int = 35
	var symB int = 64
	fmt.Scan(&x, &y, &max)

	for j := 1; j <= x; j++ {
		for k := 1; k <= y; k++ {
			for a := symA; a <= symA; a++ {
				for b := symB; b <= symB; b++ {

					str := string(rune(a)) + string(rune(b)) + strconv.Itoa(j) + strconv.Itoa(k) + string(rune(b)) + string(rune(a)) + "|"
					result += str
					count++
					symA++
					symB++

					if symA > 55 {
						symA = 35
					}

					if symB > 96 {
						symB = 64
					}
					break

				}
				break
			}
			if count == max {
				break
			}
		}

		if count == max {
			break
		}

	}

	fmt.Println(result)
}
