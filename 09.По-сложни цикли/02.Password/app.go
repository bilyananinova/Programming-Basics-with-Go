package main

import "fmt"

func main() {
	var username, password, input string

	fmt.Scanln(&username)
	fmt.Scanln(&password)

	for true {

		if password == input {
			fmt.Printf("Welcome %s!", username)
			break
		}
		fmt.Scanln(&input)
	}

}
