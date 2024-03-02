package main

import "fmt"

func main() {
	var moves, p1Count, p2Count, p3Count, p4Count, p5Count, p6Count int
	var num, score float32

	fmt.Scanln(&moves)

	for i := 0; i < moves; i++ {
		fmt.Scanln(&num)

		if num >= 0 && num <= 9 {
			score += num * 0.20
			p1Count++
		} else if num >= 10 && num <= 19 {
			p2Count++
			score += num * 0.30
		} else if num >= 20 && num <= 29 {
			p3Count++
			score += num * 0.40
		} else if num >= 30 && num <= 39 {
			p4Count++
			score += 50
		} else if num >= 40 && num <= 50 {
			p5Count++
			score += 100
		} else {
			p6Count++
			score = score / 2
		}
	}

	p1 := float32(p1Count) / float32(moves) * 100
	p2 := float32(p2Count) / float32(moves) * 100
	p3 := float32(p3Count) / float32(moves) * 100
	p4 := float32(p4Count) / float32(moves) * 100
	p5 := float32(p5Count) / float32(moves) * 100
	p6 := float32(p6Count) / float32(moves) * 100

	fmt.Printf("%.2f\n", score)
	fmt.Printf("From 0 to 9: %.2f%%\n", p1)
	fmt.Printf("From 10 to 19: %.2f%%\n", p2)
	fmt.Printf("From 20 to 29: %.2f%%\n", p3)
	fmt.Printf("From 30 to 39: %.2f%%\n", p4)
	fmt.Printf("From 40 to 50: %.2f%%\n", p5)
	fmt.Printf("Invalid numbers: %.2f%%\n", p6)

}
