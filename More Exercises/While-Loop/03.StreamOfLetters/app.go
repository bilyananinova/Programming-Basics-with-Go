package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var input, result, str string
	var c, o, n int

	for input != "End" {
		fmt.Scanln(&input)

		r, _ := utf8.DecodeRuneInString(input)

		if r >= 65 && r <= 90 || r >= 97 && r <= 122 {

			if input == "c" && c < 1 {
				c++
				continue
			} else if input == "o" && o < 1 {
				o++
				continue
			} else if input == "n" && n < 1 {
				n++
				continue
			}

			if c > 0 && o > 0 && n > 0 {
				result += str + " "
				str = ""
				c = 0
				o = 0
				n = 0
			}

			str += input
		}
	}

	fmt.Println(result)
}
