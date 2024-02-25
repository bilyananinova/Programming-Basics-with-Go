package main

import "fmt"

func main() {
	var weather string

	fmt.Scan(&weather)

	if weather == "sunny" {
		fmt.Println("It's warm outside!")
	} else {
		fmt.Println("It's cold outside!")
	}
}
