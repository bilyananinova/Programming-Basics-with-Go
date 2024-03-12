package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var firstLetter, secondLetter, thirdLetter, result string
	var count int
	fmt.Scan(&firstLetter, &secondLetter, &thirdLetter)

	start, _ := utf8.DecodeRuneInString(firstLetter)
	end, _ := utf8.DecodeRuneInString(secondLetter)

	for i := int(start); i <= int(end); i++ {
		for j := int(start); j <= int(end); j++ {
			for k := int(start); k <= int(end); k++ {

				str := string(rune(i)) + string(rune(j)) + string(rune(k)) + " "

				if !strings.Contains(str, thirdLetter) {
					result += str
					count++
				}

				str = ""

			}

		}
	}

	fmt.Printf("%s%d", result, count)
}
