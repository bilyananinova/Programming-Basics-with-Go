package main

import "fmt"

func main() {
	var restDays int
	var playLimit int = 30000

	fmt.Scanln(&restDays)

	workDays := 365 - restDays
	playTimeInWorkDays := workDays * 63
	playTimeInRestDays := restDays * 127

	playTime := playTimeInWorkDays + playTimeInRestDays

	if playTime < playLimit {
		fmt.Println("Tom sleeps well")

		h := (playLimit - playTime) / 60
		m := (playLimit - playTime) % 60

		fmt.Printf("%d hours and %d minutes less for play", h, m)
	} else {
		fmt.Println("Tom will run away")

		h := (playTime - playLimit) / 60
		m := (playTime - playLimit) % 60

		fmt.Printf("%d hours and %d minutes more for play", h, m)
	}
}
