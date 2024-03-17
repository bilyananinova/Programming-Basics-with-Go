package main

import "fmt"

func main() {
	var team, result string
	var games, score, w, d, l int

	fmt.Scanln(&team)
	fmt.Scanln(&games)

	for i := 1; i <= games; i++ {
		fmt.Scanln(&result)

		if result == "W" {
			w++
			score = score + 3
		} else if result == "D" {
			d++
			score = score + 1
		} else if result == "L" {
			l++
		}
	}

	if games == 0 {
		fmt.Printf("%s hasn't played any games during this season.", team)
	} else if games > 0 {
		fmt.Printf("%s has won %d points during this season.\n", team, score)
		fmt.Println("Total stats:")
		fmt.Printf("## W: %d\n", w)
		fmt.Printf("## D: %d\n", d)
		fmt.Printf("## L: %d\n", l)

		winsPercent := float64(w) / float64(games) * 100
		fmt.Printf("Win rate: %.2f%%", winsPercent)

	}
}
