package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var capacity int
	var input string
	var income float32
	fmt.Scanln(&capacity)

	scanner.Scan()
	input = scanner.Text()

	for input != "Movie time!" {

		people, _ := strconv.Atoi(input)

		if capacity < people {
			fmt.Println("The cinema is full.")
			break
		}

		capacity -= people

		income += float32(people) * 5.0

		if people%3 == 0 {
			income -= 5.0
		}

		scanner.Scan()
		input = scanner.Text()
	}

	if input == "Movie time!" {
		fmt.Printf("There are %d seats left in the cinema.\n", capacity)
	}

	fmt.Printf("Cinema income - %.0f lv.", income)

}
