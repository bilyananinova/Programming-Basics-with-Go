package main

import "fmt"

func main() {
	var favBook, input string
	var count int
	fmt.Scanln(&favBook)
	fmt.Scanln(&input)

	for input != "NoMoreBooks" {

		if favBook == input {
			fmt.Printf("You checked %d books and found it.", count)
			break
		}
		count++

		fmt.Scanln(&input)

	}

	if input == "NoMoreBooks" {
		fmt.Println("The book you search is not here!")
		fmt.Printf("You checked %d books.", count)
	}
}
