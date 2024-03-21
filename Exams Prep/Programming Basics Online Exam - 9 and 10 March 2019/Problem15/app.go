package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var clients, b, c, l, a, shakes, bars int
	var workout string

	fmt.Scanln(&clients)

	for i := 0; i < clients; i++ {
		scanner.Scan()
		workout = scanner.Text()

		if workout == "Back" {
			b++
		} else if workout == "Chest" {
			c++
		} else if workout == "Legs" {
			l++
		} else if workout == "Abs" {
			a++
		} else if workout == "Protein shake" {
			shakes++
		} else if workout == "Protein bar" {
			bars++
		}
	}

	workoutClients := float32(b+c+l+a) / float32(clients) * 100
	proteinClients := float32(bars+shakes) / float32(clients) * 100

	fmt.Printf("%d - back\n", b)
	fmt.Printf("%d - chest\n", c)
	fmt.Printf("%d - legs\n", l)
	fmt.Printf("%d - abs\n", a)
	fmt.Printf("%d - protein shake\n", shakes)
	fmt.Printf("%d - protein bar\n", bars)
	fmt.Printf("%.2f%% - work out\n", workoutClients)
	fmt.Printf("%.2f%% - protein", proteinClients)
}
