package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, l, thirdStart int
	var result string

	fmt.Scan(&n, &l)

	for firstNum := 1; firstNum < n; firstNum++ {
		for secondNum := 1; secondNum < n; secondNum++ {

			for firstLetter := 97; firstLetter < 97+l; firstLetter++ {
				for secondLetter := 97; secondLetter < 97+l; secondLetter++ {

					if firstNum > secondNum {
						thirdStart = firstNum + 1

					} else {
						thirdStart = secondNum + 1
					}

					for thirdNum := thirdStart; thirdNum <= n; thirdNum++ {

						str := strconv.Itoa(firstNum) + strconv.Itoa(secondNum) + string(rune(firstLetter)) + string(rune(secondLetter)) + strconv.Itoa(thirdNum)

						result += str + " "
						str = ""
					}

				}

			}
		}
	}

	fmt.Println(result)
}
