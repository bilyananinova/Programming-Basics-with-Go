package main

import "fmt"

func main() {
	var floors, rooms int

	fmt.Scan(&floors, &rooms)

	for currentfloor := floors; currentfloor >= 1; currentfloor-- {
		for currentRoom := 0; currentRoom < rooms; currentRoom++ {

			if currentfloor == floors {
				fmt.Printf("L%d%d ", currentfloor, currentRoom)
			}

			if currentfloor%2 == 0 && currentfloor != floors {
				fmt.Printf("O%d%d ", currentfloor, currentRoom)
			} else if currentfloor%2 == 1 && currentfloor != floors {
				fmt.Printf("A%d%d ", currentfloor, currentRoom)
			}
		}

		fmt.Println()
	}
}
