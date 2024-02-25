package main

import "fmt"

func main() {
	var p1, p2 int
	var v, h float32

	fmt.Scan(&v, &p1, &p2, &h)

	firstPipe := float32(p1) * h
	secondPipe := float32(p2) * h
	total := firstPipe + secondPipe

	if total <= v {
		full := total / v * 100
		firstPipeFull := firstPipe / total * 100
		secondPipeFull := secondPipe / total * 100
		fmt.Printf("The pool is %.2f%% full. Pipe 1: %.2f%%. Pipe 2: %.2f%%.", full, firstPipeFull, secondPipeFull)
	} else {
		fmt.Printf("For %.2f hours the pool overflows with %.2f liters.", h, total-v)
	}

}
