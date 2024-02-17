package main

import "fmt"

func main() {
	var startScore int
	var bonus float64

	fmt.Scanln(&startScore)

	if startScore <= 100 {
		bonus = 5
	} else if startScore > 100 && startScore <= 1000 {
		bonus = float64(startScore) * 0.20
	} else if startScore > 1000 {
		bonus = float64(startScore) * 0.10
	}

	if startScore%2 == 0 {
		bonus += 1
	} else if startScore%10 == 5 {
		bonus += 2
	}

	result := float64(startScore) + bonus

	fmt.Println(bonus)
	fmt.Println(result)
}
