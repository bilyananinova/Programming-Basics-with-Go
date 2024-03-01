package main

import "fmt"

func main() {
	var n, currentNum, p1, p2, p3, p4, p5 int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&currentNum)

		if currentNum < 200 {
			p1++
		} else if currentNum >= 200 && currentNum <= 399 {
			p2++
		} else if currentNum >= 400 && currentNum <= 599 {
			p3++
		} else if currentNum >= 600 && currentNum <= 799 {
			p4++
		} else if currentNum >= 800 {
			p5++
		}

	}

	resultP1 := float32(p1) / float32(n) * 100
	resultP2 := float32(p2) / float32(n) * 100
	resultP3 := float32(p3) / float32(n) * 100
	resultP4 := float32(p4) / float32(n) * 100
	resultP5 := float32(p5) / float32(n) * 100

	fmt.Printf("%.2f%%\n", resultP1)
	fmt.Printf("%.2f%%\n", resultP2)
	fmt.Printf("%.2f%%\n", resultP3)
	fmt.Printf("%.2f%%\n", resultP4)
	fmt.Printf("%.2f%%\n", resultP5)
}
