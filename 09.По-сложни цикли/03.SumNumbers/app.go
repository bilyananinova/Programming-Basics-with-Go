package main

import "fmt"

func main() {
	var num, input, sum int
	fmt.Scanln(&num)
	fmt.Scanln(&input)

	sum += input

	for sum < num {

		fmt.Scanln(&input)
		sum += input
	}

	fmt.Println(sum)
}
