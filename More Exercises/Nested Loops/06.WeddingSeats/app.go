package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var lastSector string
	var rowsFirstSecotor, oddRowsSeats, seatsCount int
	fmt.Scanln(&lastSector)
	fmt.Scanln(&rowsFirstSecotor)
	fmt.Scanln(&oddRowsSeats)

	end, _ := utf8.DecodeRuneInString(lastSector)

	for i := 65; i <= int(end); i++ {
		for r := 1; r <= rowsFirstSecotor; r++ {

			if r%2 == 0 {
				seats := oddRowsSeats + 2

				for s := 0; s < seats; s++ {
					str := string(rune(i)) + strconv.Itoa(r) + strings.ToLower(string(rune(65+s)))
					fmt.Println(str)
					seatsCount++
				}
			} else {
				for s := 0; s < oddRowsSeats; s++ {
					str := string(rune(i)) + strconv.Itoa(r) + strings.ToLower(string(rune(65+s)))
					fmt.Println(str)
					seatsCount++
				}
			}

		}

		rowsFirstSecotor++

	}

	fmt.Println(seatsCount)
}
