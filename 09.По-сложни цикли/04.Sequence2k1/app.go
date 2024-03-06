package main

import "fmt"

func main() {
	var n, currentNum int
	currentNum = 1

	fmt.Scanln(&n)

	for currentNum <= n {

		fmt.Println(currentNum)

		currentNum = currentNum*2 + 1
	}
}
