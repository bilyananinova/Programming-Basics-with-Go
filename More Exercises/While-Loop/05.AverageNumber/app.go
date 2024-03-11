package main

import "fmt"

func main() {
	var n, currentNum int
	var sum, average float32

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&currentNum)
		sum += float32(currentNum)
	}

	average = sum / float32(n)
	fmt.Printf("%.2f", average)
}
