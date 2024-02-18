package main

import "fmt"

func main() {
	var firstTime, secondTime, thirdTime int

	fmt.Scanln(&firstTime)
	fmt.Scanln(&secondTime)
	fmt.Scanln(&thirdTime)

	var sumTime = firstTime + secondTime + thirdTime

	var minutes = sumTime / 60
	var seconds = sumTime % 60

	if seconds >= 10 {
		fmt.Printf("%d:%d", minutes, seconds)
	} else {
		fmt.Printf("%d:0%d", minutes, seconds)
	}
}
