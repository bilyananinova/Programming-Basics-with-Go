package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mens, ladys, tables int
	var result string

	fmt.Scanln(&mens)
	fmt.Scanln(&ladys)
	fmt.Scanln(&tables)

	for m := 1; m <= mens; m++ {
		for l := 1; l <= ladys; l++ {

			if tables > 0 {
				result += "(" + strconv.Itoa(m) + " <-> " + strconv.Itoa(l) + ") "
				tables--
			}
		}
	}

	fmt.Println(result)
}
