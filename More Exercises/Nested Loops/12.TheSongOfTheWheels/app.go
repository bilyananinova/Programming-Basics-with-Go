package main

import (
	"fmt"
	"strconv"
)

func main() {
	var m, counter int
	var passwords, password string

	fmt.Scanln(&m)

	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			for c := 1; c <= 9; c++ {
				for d := 1; d <= 9; d++ {

					if a < b && c > d {
						if a*b+c*d == m {
							counter++
							str := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c) + strconv.Itoa(d) + " "
							passwords += str

							if counter == 4 {
								password = str
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("%s\n", passwords)

	if password == "" {
		fmt.Println("No!")
	} else {
		fmt.Printf("Password: %s", password)
	}
}
