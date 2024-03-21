package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstResult, secondResult, thirdResult string
	var w, l, d int

	fmt.Scanln(&firstResult)
	fmt.Scanln(&secondResult)
	fmt.Scanln(&thirdResult)

	firstGame := strings.Split(firstResult, ":")
	secondGame := strings.Split(secondResult, ":")
	thirdGame := strings.Split(thirdResult, ":")

	if firstGame[0] > firstGame[1] {
		w++
	} else if firstGame[0] < firstGame[1] {
		l++
	} else if firstGame[0] == firstGame[1] {
		d++
	}

	if secondGame[0] > secondGame[1] {
		w++
	} else if secondGame[0] < secondGame[1] {
		l++
	} else if secondGame[0] == secondGame[1] {
		d++
	}

	if thirdGame[0] > thirdGame[1] {
		w++
	} else if thirdGame[0] < thirdGame[1] {
		l++
	} else if thirdGame[0] == thirdGame[1] {
		d++
	}

	fmt.Printf("Team won %d games.\n", w)
	fmt.Printf("Team lost %d games.\n", l)
	fmt.Printf("Drawn games: %d", d)
}
