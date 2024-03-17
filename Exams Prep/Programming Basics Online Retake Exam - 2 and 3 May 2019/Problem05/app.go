package main

import "fmt"

func main() {
	var n, currentNumber int
	var p1, p2, p3 float32

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&currentNumber)

		if currentNumber%2 == 0 {
			p1++
		}

		if currentNumber%3 == 0 {
			p2++
		}

		if currentNumber%4 == 0 {
			p3++
		}
	}

	fmt.Printf("%.2f%%\n", p1/float32(n)*100)
	fmt.Printf("%.2f%%\n", p2/float32(n)*100)
	fmt.Printf("%.2f%%", p3/float32(n)*100)
}
